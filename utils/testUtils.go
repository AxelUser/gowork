package utils

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/AxelUser/gowork/errors/normalizerErrors"
	"github.com/AxelUser/gowork/models/configs"
	"github.com/AxelUser/gowork/models/dataModels"
)

func CreateRawData(aliases []string, countPerSkill int, createUnique bool) map[string][]dataModels.VacancyStats {
	statsMap := make(map[string][]dataModels.VacancyStats)
	for i, alias := range aliases {
		var stats []dataModels.VacancyStats
		for j := 0; j < countPerSkill; j++ {
			salaryFrom := float32((i + 1) * 10000)
			salaryTo := float32((i + 1) * 20000)

			var id string
			if createUnique {
				id = strconv.Itoa((i+1)*100000 + j)
			} else {
				id = "1"
			}
			s := dataModels.NewVacancyStats(id, "test.com", &salaryFrom, &salaryTo, "RUB", alias)
			stats = append(stats, s)
		}
		statsMap[alias] = stats
	}
	return statsMap
}

func CreateOntology(aliases []string, emptyRules bool, addRuleForItself bool) []configs.OntologyData {
	var ontology []configs.OntologyData
	for _, alias := range aliases {
		o := configs.OntologyData{Alias: alias, Caption: alias}
		if !emptyRules {
			o.Rules = make(map[string]float32)
			for _, skill := range aliases {
				if skill == alias {
					if addRuleForItself {
						o.Rules[skill] = 1
					}
				} else {
					o.Rules[skill] = 0.1
				}
			}
		}
		ontology = append(ontology, o)
	}

	return ontology
}

func CheckNormalizerErrorCode(errs []error, code int) error {
	if len(errs) > 0 {
		for _, err := range errs {
			switch err.(type) {
			case normalizerErrors.NormalizerError:
				e := err.(normalizerErrors.NormalizerError)
				if e.CaseCode == code {
					return nil
				}
				return fmt.Errorf("NormalizerError is must be with code %d: %d", code, e.CaseCode)
			}
		}
		return fmt.Errorf("Error is not NormalizerError: %s", reflect.TypeOf(errs[0]))
	}
	return fmt.Errorf("Expected error")
}
