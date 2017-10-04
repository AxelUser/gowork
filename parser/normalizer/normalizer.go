package normalizer

import (
	"errors"
	"strings"

	"github.com/AxelUser/gowork/models"
)

func checkRawData(ontologyInfos []models.OntologyData, rawData map[string][]models.VacancyStats) error {
	var missingData []string
	var missingRules []string
	errMsg := ""

	for _, info := range ontologyInfos {
		if len(rawData[info.Alias]) == 0 {
			missingData = append(missingData, info.Alias)
		}
		if len(info.Rules) == 0 {
			missingRules = append(missingRules, info.Alias)
		}
	}

	if len(missingData) > 0 {
		errMsg += "Missing data: " + strings.Join(missingData, ", ") + ". "
	}

	if len(missingRules) > 0 {
		errMsg += "Missing rules: " + strings.Join(missingRules, ", ") + ". "
	}

	if errMsg != "" {
		return errors.New("Normalizing failed. " + errMsg)
	}

	return nil
}

// NormalizeRawData proceeds vacancies and normalize them for training set
func NormalizeRawData(ontologyInfos []models.OntologyData, rawData map[string][]models.VacancyStats) (map[string][]int, error) {
	err := checkRawData(ontologyInfos, rawData)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
