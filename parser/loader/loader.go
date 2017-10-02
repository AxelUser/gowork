package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/AxelUser/gowork/errors"
	"github.com/AxelUser/gowork/events"
	"github.com/AxelUser/gowork/models"
)

func createBaseURL(config models.ParserConfig) (string, error) {
	req, err := http.NewRequest("GET", config.URL, nil)
	if err != nil {
		return "", errors.NewLoadDataError(config.URL, "Could not create request for Base URL", err)
	}

	q := req.URL.Query()
	for k, v := range config.Defaults {
		q.Add(k, fmt.Sprintf("%v", v))
	}

	req.URL.RawQuery = q.Encode()
	url := req.URL.String()
	return url, nil
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

func loadDataPerSkill(alias string, url string) ([]models.VacancyStats, error) {
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

	data := make([]models.VacancyStats, len(page.Items))
	// Handle empty items and log!
	for i, v := range page.Items {
		data[i] = models.NewVacancyStats(v.ID, v.URL, v.Salary.From, v.Salary.To, v.Salary.Currency)
	}
	return data, nil
}

func loadDataPerSkillAsync(jobsCh <-chan models.LoaderJob, eventCh chan<- events.DataLoadedEvent, errCh chan<- error, wg *sync.WaitGroup) {
	for job := range jobsCh {
		stats, err := loadDataPerSkill(job.Alias, job.URL)
		if err != nil {
			errCh <- err
		} else {
			eventCh <- events.NewDataLoadedEvent(job.Alias, job.URL, stats)
		}
	}

	wg.Done()
}

func loadAll(urls map[string]string, count int) ([]models.VacancyStats, error) {
	var all []models.VacancyStats
	dataCh := make(chan events.DataLoadedEvent, len(urls))
	errCh := make(chan error, len(urls))
	jobsCh := make(chan models.LoaderJob, len(urls))
	var wg sync.WaitGroup

	wg.Add(count)
	//create workers for loading vacancies
	for i := 0; i < count; i++ {
		go loadDataPerSkillAsync(jobsCh, dataCh, errCh, &wg)
	}

	log.Printf("Workers in pool: %d\n", count)

	for alias, url := range urls {
		jobsCh <- models.NewLoaderJob(alias, url)
	}
	close(jobsCh)

	for i := 0; i < len(urls); i++ {
		event := <-dataCh
		log.Printf(event.String())
		all = append(all, event.Data...)
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
