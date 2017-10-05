package normalizer

import (
	nErrors "github.com/AxelUser/gowork/errors/normalizerErrors"
	"github.com/AxelUser/gowork/models"
)

func checkRawData(ontologyInfos []models.OntologyData, rawData map[string][]models.VacancyStats) []error {
	var missingData []string
	var missingRules []string

	var checkingErrors []error

	for _, info := range ontologyInfos {
		if len(rawData[info.Alias]) == 0 {
			missingData = append(missingData, info.Alias)
		}
		if len(info.Rules) == 0 {
			missingRules = append(missingRules, info.Alias)
		}
	}

	if len(missingData) > 0 {
		checkingErrors = append(checkingErrors, nErrors.New(nErrors.CaseCodeMissingData, missingData, nil))
	}

	if len(missingRules) > 0 {
		checkingErrors = append(checkingErrors, nErrors.New(nErrors.CaseCodeEmptyRules, missingRules, nil))
	}

	return checkingErrors
}

func getPlainData(groupedData map[string][]models.VacancyStats) []models.VacancyStats {
	var plainData []models.VacancyStats

	for _, group := range groupedData {
		plainData = append(plainData, group...)
	}

	return plainData
}

func resolveDublicates(ontologyInfos []models.OntologyData, plainRawData []models.VacancyStats) ([]models.VacancyStats, int) {
	uniqueStatsFrequencyMap := make(map[string]int)
	totalDublicates := 0
	var uniqueData []models.VacancyStats

	for _, stat := range plainRawData {
		if _, ok := uniqueStatsFrequencyMap[stat.ID]; ok {
			uniqueStatsFrequencyMap[stat.ID]++
			totalDublicates++
		} else {
			uniqueStatsFrequencyMap[stat.ID] = 1
			uniqueData = append(uniqueData, stat)
		}
	}

	return uniqueData, totalDublicates
}

// NormalizeRawData proceeds vacancies and normalize them for training set
func NormalizeRawData(ontologyInfos []models.OntologyData, rawData map[string][]models.VacancyStats) (map[string][]int, []error) {
	errs := checkRawData(ontologyInfos, rawData)
	if len(errs) > 0 {
		return nil, errs
	}

	return nil, nil
}
