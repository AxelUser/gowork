package currencyLoader

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/AxelUser/gowork/models/dataModels"

	"github.com/AxelUser/gowork/errors"
	"github.com/AxelUser/gowork/models/api"
)

const currencyAPIBaseURL string = "http://api.fixer.io/latest?base="
const baseCurrency string = "RUB"

var loadedRates *api.CurrencyRates

// Load is for loading latest currency rates
func Load() (*api.CurrencyRates, error) {
	resp, err := http.Get(currencyAPIBaseURL + baseCurrency)
	if err != nil {
		return nil, errors.NewLoadDataError(currencyAPIBaseURL, "", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewLoadDataError(currencyAPIBaseURL, "", err)
	}

	var rates api.CurrencyRates
	err = json.Unmarshal(body, &rates)
	if err != nil {
		return nil, errors.NewLoadDataError(currencyAPIBaseURL, "", err)
	}

	if rates.Error != "" {
		return nil, errors.NewLoadDataError(currencyAPIBaseURL, rates.Error, nil)
	}

	loadedRates = &rates
	return &rates, nil
}

// ConvertSalary returns VacancyStats with converted salary
func ConvertSalary(stat dataModels.VacancyStats) dataModels.VacancyStats {
	if stat.Currency == baseCurrency {
		return stat
	}

	curRate := loadedRates.Rates[stat.Currency]

	if stat.SalaryFrom != nil {
		newSalary := *stat.SalaryFrom / curRate
		stat.SalaryFrom = &newSalary
	}

	if stat.SalaryTo != nil {
		newSalary := *stat.SalaryTo / curRate
		stat.SalaryTo = &newSalary
	}

	return stat
}
