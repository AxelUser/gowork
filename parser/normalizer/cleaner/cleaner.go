package cleaner

import (
	nErrors "github.com/AxelUser/gowork/errors/normalizerErrors"
	"github.com/AxelUser/gowork/models/configs"
	"github.com/AxelUser/gowork/models/dataModels"
)

// Clean is a function to check raw data and remove all dublicates
func Clean(ontologyInfos []configs.OntologyData, rawData map[string][]dataModels.VacancyStats) ([]dataModels.VacancyStats, []error) {
	errs := checkRawData(ontologyInfos, rawData)
	if len(errs) > 0 {
		return nil, errs
	}

	plainData := getPlainData(rawData)

	uniqueData, _ := resolveDublicates(plainData)

	return uniqueData, nil
}

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
			if _, ok := info.Rules[info.Alias]; !ok {
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
			dataUpd := uniqueStatsMap[stat.ID]
			dataUpd.AddSkill(stat.Skills[0])
			uniqueStatsMap[stat.ID] = dataUpd
			totalDublicates++
		} else {
			uniqueStatsMap[stat.ID] = stat
		}
	}

	for _, stat := range uniqueStatsMap {
		uniqueData = append(uniqueData, stat)
	}

	return uniqueData, totalDublicates
}
