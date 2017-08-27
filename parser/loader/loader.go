package loader

import (
	"encoding/json"
	"gowork/errors"
	"gowork/models"
	"io/ioutil"
	"net/http"
)

func loadDataPerSkill(alias string, url string) (*[]models.VacancyStats, error) {
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
	jsonErr := json.Unmarshal(body, page)
	if jsonErr != nil {
		return nil, errors.NewLoadSkillError(alias, "Could not unmarshal JSON", jsonErr)
	}

	data := make([]models.VacancyStats, len(page.Items))
	// Handle empty items and log!
	for i, v := range page.Items {
		data[i] = models.NewVacancyStats(v.ID, v.URL, v.Salary.From, v.Salary.To, v.Salary.Currency)
	}
	return &data, nil
}

func createBaseURL(config models.ParserConfig) (*string, error) {
	req, err := http.NewRequest("GET", config.URL, nil)
	if err != nil {
		return nil, errors.NewLoadDataError(config.URL, "Could not create request for Base URL", err)
	}

	q := req.URL.Query()
	for k, v := range config.Defaults {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()
	url := req.URL.String()
	return &url, nil
}

func createURLs(baseURL string, queries map[string]string) (*map[string]string, error) {
	skillURLMap := make(map[string]string)

	for alias, text := range queries {
		req, err := http.NewRequest("GET", baseURL, nil)
		if err != nil {
			return nil, errors.NewLoadSkillError(alias, "Could not create request for skill", err)
		}

		q := req.URL.Query()
		q.Add("text", text)

		req.URL.RawQuery = q.Encode()

		skillURL := req.URL.String()

		skillURLMap[alias] = skillURL
	}

	return &skillURLMap, nil
}

// Load data from HeadHunter API
func Load(models.ParserConfig) []models.VacancyStats {
	var allStats []models.VacancyStats
	return allStats
}
