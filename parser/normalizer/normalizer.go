package normalizer

import (
	nErrors "github.com/AxelUser/gowork/errors/normalizerErrors"
	"github.com/AxelUser/gowork/models/configs"
	"github.com/AxelUser/gowork/models/dataModels"
)

func checkRawData(ontologyInfos []configs.OntologyData, rawData map[string][]dataModels.VacancyStats) []error {
	var missingData []string
	var missingRules []string
	var missingRulesForThemselves []string

	var checkingErrors []error

	for _, info := range ontologyInfos {
		if len(rawData[info.Alias]) == 0 {
			missingData = append(missingData, info.Alias)
		}
		if len(info.Rules) == 0 {
			missingRules = append(missingRules, info.Alias)
			missingRulesForThemselves = append(missingRulesForThemselves, info.Alias)
		} else {
			for skill := range info.Rules {
				if skill == info.Alias {
					continue
				}
				missingRulesForThemselves = append(missingRulesForThemselves, info.Alias)
			}
		}
	}

	if len(missingData) > 0 {
		checkingErrors = append(checkingErrors, nErrors.New(nErrors.CaseCodeMissingData, missingData, nil))
	}

	if len(missingRules) > 0 {
		checkingErrors = append(checkingErrors, nErrors.New(nErrors.CaseCodeEmptyRules, missingRules, nil))
	}

	if len(missingRulesForThemselves) > 0 {
		checkingErrors = append(checkingErrors,
			nErrors.New(nErrors.CaseCodeOntologyMissingRuleForSameSkill, missingRulesForThemselves, nil))
	}

	return checkingErrors
}

func getPlainData(groupedData map[string][]dataModels.VacancyStats) []dataModels.VacancyStats {
	var plainData []dataModels.VacancyStats

	for _, group := range groupedData {
		plainData = append(plainData, group...)
	}

	return plainData
}

func resolveDublicates(plainRawData []dataModels.VacancyStats) ([]dataModels.VacancyStats, int) {
	uniqueStatsMap := make(map[string]dataModels.VacancyStats)
	totalDublicates := 0
	var uniqueData []dataModels.VacancyStats

	for _, stat := range plainRawData {
		if _, ok := uniqueStatsMap[stat.ID]; ok {
			// Adds skill for what it was loaded
			uniqueStatsMap[stat.ID].AddSkill(stat.Skills[0])
			totalDublicates++
		} else {
			uniqueStatsMap[stat.ID] = stat
			uniqueData = append(uniqueData, stat)
		}
	}

	return uniqueData, totalDublicates
}

func normalizeInputsAndOutputs(ontology []configs.OntologyData, data []dataModels.VacancyStats) ([]float32, []float32) {
	return nil, nil
}

// NormalizeRawData proceeds vacancies and normalize them for training set
func NormalizeRawData(ontologyInfos []configs.OntologyData, rawData map[string][]dataModels.VacancyStats) (map[string][]int, []error) {
	errs := checkRawData(ontologyInfos, rawData)
	if len(errs) > 0 {
		return nil, errs
	}

	plainData := getPlainData(rawData)

	resolveDublicates(plainData)

	return nil, nil
}
