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

// BaseCurrency is currency for which rates are loaded
const BaseCurrency string = "RUB"

var baseCurrencyCodes = []string{"RUB", "RUR"}

// Load is for loading latest currency rates
func Load() (*api.CurrencyRates, error) {
	resp, err := http.Get(currencyAPIBaseURL + BaseCurrency)
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

	return &rates, nil
}

// ConvertSalary returns VacancyStats with converted salary
func ConvertSalary(loadedRates api.CurrencyRates, stat dataModels.VacancyStats) dataModels.VacancyStats {
	if !IsForeignCurrency(stat) {
		return stat
	}

	curRate := loadedRates.Rates[stat.Currency]
	var newSalaryFrom *float32
	var newSalaryTo *float32

	if stat.SalaryFrom != nil {
		temp := *stat.SalaryFrom / curRate
		newSalaryFrom = &temp
	}

	if stat.SalaryTo != nil {
		temp := *stat.SalaryTo / curRate
		newSalaryTo = &temp
	}

	return dataModels.NewVacancyStats(stat.ID, stat.URL, newSalaryFrom, newSalaryTo, BaseCurrency)
}

// IsForeignCurrency check if VacancyStats has salary in foreign currency
func IsForeignCurrency(stat dataModels.VacancyStats) bool {
	for _, code := range baseCurrencyCodes {
		if stat.Currency == code {
			return false
		}
	}
	return true
}
