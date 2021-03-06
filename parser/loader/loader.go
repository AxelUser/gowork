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
	"time"

	"github.com/AxelUser/gowork/parser/loader/currencyLoader"

	"github.com/AxelUser/gowork/errors"
	"github.com/AxelUser/gowork/events"
	"github.com/AxelUser/gowork/models/api"
	"github.com/AxelUser/gowork/models/configs"
	"github.com/AxelUser/gowork/models/dataModels"
)

// Load data from HeadHunter API
func Load(config configs.ParserConfig) (map[string][]dataModels.VacancyStats, error) {

	allStats, err := loadVacancyStatsForAllSkills(config)
	if err != nil {
		return allStats, err
	}

	convertedStats, err := convertSalariesToBaseCurrency(allStats)
	if err != nil {
		return convertedStats, err
	}

	return convertedStats, nil
}

func createBaseURL(config configs.ParserConfig) (string, error) {
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

func createURLs(baseURL string, queries []configs.ParserQuery) (map[string]string, error) {
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

func loadPage(alias string, url string) (*api.VacancySearchPage, error) {
	res, httpErr := http.Get(url)
	if httpErr != nil {
		return nil, errors.NewLoadSkillError(alias, "Could not send GET request", httpErr)
	}
	defer res.Body.Close()

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return nil, errors.NewLoadSkillError(alias, "Could not read data from Body", bodyErr)
	}

	var page api.VacancySearchPage
	jsonErr := json.Unmarshal(body, &page)
	if jsonErr != nil {
		return nil, errors.NewLoadSkillError(alias, "Could not unmarshal JSON", jsonErr)
	}

	if len(page.Errors) > 0 {
		apiErr := page.Errors[0]
		return nil, errors.NewLoadSkillError(alias, apiErr.Type+" <"+apiErr.Value+">", nil)
	}

	return &page, nil
}

func loadAllPages(alias string, firstPageURL string, firstPage api.VacancySearchPage) ([]api.VacancySearchPage, error) {
	pages := []api.VacancySearchPage{firstPage}
	pageURL := firstPageURL
	currentPage := &firstPage

	for !isLastPage(*currentPage) {
		pageURL, err := getNextPageURL(alias, pageURL, currentPage.Page+1)
		if err != nil {
			return pages, err
		}

		currentPage, err = loadPage(alias, pageURL)
		if err != nil {
			return pages, err
		}

		pages = append(pages, *currentPage)
	}

	return pages, nil
}

func parseVacancyStats(alias string, page api.VacancySearchPage) []dataModels.VacancyStats {
	data := make([]dataModels.VacancyStats, len(page.Items))
	// Handle empty items and log!
	for i, v := range page.Items {
		data[i] = dataModels.NewVacancyStats(v.ID, v.URL, v.Salary.From, v.Salary.To, v.Salary.Currency, alias)
	}

	return data
}

func isLastPage(page api.VacancySearchPage) bool {
	return page.Page >= page.Pages-1
}

func loadDataPerSkillAsync(jobsCh <-chan dataModels.LoaderJob, eventCh chan<- events.DataLoadedEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobsCh {
		pageURL := job.URL
		pageModel, err := loadPage(job.Alias, pageURL)
		if err != nil {
			eventCh <- events.NewDataLoadedEventWithError(job.Alias, job.URL, nil, err)
		} else {
			pages, err := loadAllPages(job.Alias, pageURL, *pageModel)
			var allStats []dataModels.VacancyStats
			for _, page := range pages {
				allStats = append(allStats, parseVacancyStats(job.Alias, page)...)
			}

			eventCh <- events.NewDataLoadedEventWithError(job.Alias, pageURL, allStats, err)
		}
	}
}

func loadAll(urls map[string]string, count int) (map[string][]dataModels.VacancyStats, int, error) {
	all := make(map[string][]dataModels.VacancyStats)
	dataCh := make(chan events.DataLoadedEvent, len(urls))
	jobsCh := make(chan dataModels.LoaderJob, len(urls))
	var wg sync.WaitGroup

	wg.Add(count)
	//create workers for loading vacancies
	for i := 0; i < count; i++ {
		go loadDataPerSkillAsync(jobsCh, dataCh, &wg)
	}

	log.Printf("Workers in pool: %d\n", count)

	for alias, url := range urls {
		jobsCh <- dataModels.NewLoaderJob(alias, url)
	}
	close(jobsCh)

	totalCount := 0
	for i := 0; i < len(urls); i++ {
		event := <-dataCh
		log.Println(event)
		if event.HasData() {
			totalCount += len(event.Data)
			all[event.Skill] = event.Data
		}
	}

	log.Println("Waiting for workers to complete their jobs")
	wg.Wait()

	return all, totalCount, nil
}

func loadVacancyStatsForAllSkills(config configs.ParserConfig) (map[string][]dataModels.VacancyStats, error) {
	timeStart := time.Now()

	allStats := make(map[string][]dataModels.VacancyStats)

	baseURL, err := createBaseURL(config)
	if err != nil {
		return allStats, err
	}
	urls, err := createURLs(baseURL, config.Queries)
	if err != nil {
		return allStats, err
	}
	log.Printf("Loading vacancies from %s", config.URL)

	allStats, totalCount, err := loadAll(urls, config.WorkersCount)
	if err != nil {
		return nil, err
	}

	elapsed := time.Since(timeStart)
	log.Printf("Loaded %d item(s) in %s\n", totalCount, elapsed)

	return allStats, nil
}

func convertSalariesToBaseCurrency(allStats map[string][]dataModels.VacancyStats) (map[string][]dataModels.VacancyStats, error) {
	log.Printf("Loading latest currency rates for %s ", currencyLoader.BaseCurrency)
	rates, err := currencyLoader.Load()
	if err != nil {
		return allStats, err
	}
	log.Printf("Loaded rates for %s", rates.Date)

	log.Println("Converting vacancies with foreign currencies")
	totalConvertions := 0
	for alias, group := range allStats {
		for i, stat := range group {
			if currencyLoader.IsForeignCurrency(stat) {
				allStats[alias][i] = currencyLoader.ConvertSalary(*rates, stat)
				totalConvertions++
			}
		}
	}
	log.Printf("Converted %d item(s)", totalConvertions)

	return allStats, nil
}
