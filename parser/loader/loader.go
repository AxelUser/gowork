package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	urlModule "net/url"
	"strconv"
	"sync"

	"github.com/AxelUser/gowork/errors"
	"github.com/AxelUser/gowork/events"
	"github.com/AxelUser/gowork/models"
)

func createBaseURL(config models.ParserConfig) (string, error) {
	url, err := urlModule.Parse(config.URL)
	if err != nil {
		return "", errors.NewLoadDataError(config.URL, "Could not parse Base URL", err)
	}

	q := url.Query()
	for k, v := range config.Defaults {
		q.Add(k, fmt.Sprintf("%v", v))
	}

	url.RawQuery = q.Encode()
	urlString := url.String()
	return urlString, nil
}

func createURLs(baseURL string, queries []models.ParserQuery) (map[string]string, error) {
	skillURLMap := make(map[string]string)

	for _, query := range queries {
		req, err := http.NewRequest("GET", baseURL, nil)
		if err != nil {
			return nil, errors.NewLoadSkillError(query.Alias, "Could not create request for skill", err)
		}

		q := req.URL.Query()
		q.Add("text", query.Text)

		req.URL.RawQuery = q.Encode()

		skillURL := req.URL.String()

		skillURLMap[query.Alias] = skillURL
	}

	return skillURLMap, nil
}

func getNextPageURL(alias string, url string, nextPage int) (string, error) {
	urlBuild, err := urlModule.Parse(url)
	if err != nil {
		return "", errors.NewLoadSkillError(alias, "Could not create URL for next page", err)
	}

	q := urlBuild.Query()
	q.Set("page", strconv.Itoa(nextPage))
	urlBuild.RawQuery = q.Encode()
	nextPageURL := urlBuild.String()

	return nextPageURL, nil
}

func loadPage(alias string, url string) (*models.VacancySearchPage, error) {
	res, httpErr := http.Get(url)
	if httpErr != nil {
		return nil, errors.NewLoadSkillError(alias, "Could not send GET request", httpErr)
	}
	defer res.Body.Close()

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return nil, errors.NewLoadSkillError(alias, "Could not read data from Body", bodyErr)
	}

	var page models.VacancySearchPage
	jsonErr := json.Unmarshal(body, &page)
	if jsonErr != nil {
		return nil, errors.NewLoadSkillError(alias, "Could not unmarshal JSON", jsonErr)
	}

	return &page, nil
}

func parseVacancyStats(page models.VacancySearchPage) []models.VacancyStats {
	data := make([]models.VacancyStats, len(page.Items))
	// Handle empty items and log!
	for i, v := range page.Items {
		data[i] = models.NewVacancyStats(v.ID, v.URL, v.Salary.From, v.Salary.To, v.Salary.Currency)
	}

	return data
}

func isLastPage(page models.VacancySearchPage) bool {
	return page.Page >= page.Pages-1
}

func loadDataPerSkillAsync(jobsCh <-chan models.LoaderJob, eventCh chan<- events.DataLoadedEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobsCh {
		pageURL := job.URL
		var pages []models.VacancySearchPage
		pageModel, err := loadPage(job.Alias, pageURL)
		if err != nil {
			eventCh <- events.NewDataLoadedEventWithError(job.Alias, job.URL, err)
		} else {
			pages = append(pages, *pageModel)
			var allStats []models.VacancyStats

			for !isLastPage(*pageModel) {
				pageURL, err = getNextPageURL(job.Alias, pageURL, pageModel.Page+1)
				if err != nil {
					eventCh <- events.NewDataLoadedEventWithError(job.Alias, job.URL, err)
					return
				}

				pageModel, err = loadPage(job.Alias, pageURL)
				if err != nil {
					eventCh <- events.NewDataLoadedEventWithError(job.Alias, job.URL, err)
					return
				}

				pages = append(pages, *pageModel)
			}

			for _, page := range pages {
				allStats = append(allStats, parseVacancyStats(page)...)
			}
			eventCh <- events.NewDataLoadedEvent(job.Alias, job.URL, allStats)
		}
	}

}

func loadAll(urls map[string]string, count int) ([]models.VacancyStats, error) {
	var all []models.VacancyStats
	dataCh := make(chan events.DataLoadedEvent, len(urls))
	jobsCh := make(chan models.LoaderJob, len(urls))
	var wg sync.WaitGroup

	wg.Add(count)
	//create workers for loading vacancies
	for i := 0; i < count; i++ {
		go loadDataPerSkillAsync(jobsCh, dataCh, &wg)
	}

	log.Printf("Workers in pool: %d\n", count)

	for alias, url := range urls {
		jobsCh <- models.NewLoaderJob(alias, url)
	}
	close(jobsCh)

	for i := 0; i < len(urls); i++ {
		event := <-dataCh
		log.Println(event)
		if event.IsSuccess() {
			all = append(all, event.Data...)
		}
	}
	wg.Wait()
	return all, nil
}

// Load data from HeadHunter API
func Load(config models.ParserConfig) ([]models.VacancyStats, error) {
	var allStats []models.VacancyStats

	baseURL, err := createBaseURL(config)
	if err != nil {
		return allStats, err
	}
	urls, err := createURLs(baseURL, config.Queries)
	if err != nil {
		return allStats, err
	}
	log.Printf("Loading vacancies from %s", config.URL)

	allStats, err = loadAll(urls, config.WorkersCount)
	if err != nil {
		return nil, err
	}
	log.Printf("\nLoaded %d item(s)\n", len(allStats))

	return allStats, nil
}
