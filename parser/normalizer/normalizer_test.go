package normalizer

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/AxelUser/gowork/errors/normalizerErrors"
	"github.com/AxelUser/gowork/models"
)

func createRawData(aliases []string, countPerSkill int, createUnique bool) map[string][]models.VacancyStats {
	statsMap := make(map[string][]models.VacancyStats)
	for i, alias := range aliases {
		var stats []models.VacancyStats
		for j := 0; i < countPerSkill; i++ {
			salaryFrom := (i + 1) * 10000
			salaryTo := (i + 1) * 20000

			var id string
			if createUnique {
				id = strconv.Itoa((i+1)*100000 + j)
			} else {
				id = "1"
			}
			s := models.NewVacancyStats(id, "test.com", &salaryFrom, &salaryTo, "RUB")
			stats = append(stats, s)
		}
		statsMap[alias] = stats
	}
	return statsMap
}

func createOntology(aliases []string, emptyRules bool, addRuleForItself bool) []models.OntologyData {
	var ontology []models.OntologyData
	for _, alias := range aliases {
		o := models.OntologyData{Alias: alias, Caption: alias}
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

func checkNormalizerErrorCode(errs []error, code int) error {
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

func TestCheckRawData_NoData_ReturnsError(t *testing.T) {
	raw := createRawData([]string{"js", "css"}, 10, true)
	ontology := createOntology([]string{"js", "css", "html"}, false, true)

	errs := checkRawData(ontology, raw)

	err := checkNormalizerErrorCode(errs, normalizerErrors.CaseCodeMissingData)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckRawData_EmptyRules_ReturnsErrors(t *testing.T) {
	raw := createRawData([]string{"js", "css", "html"}, 10, true)
	ontology := createOntology([]string{"js", "css", "html"}, true, false)

	errs := checkRawData(ontology, raw)

	err := checkNormalizerErrorCode(errs, normalizerErrors.CaseCodeEmptyRules)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckRawData_MissingRulesForSameSkill_ReturnsError(t *testing.T) {
	raw := createRawData([]string{"js", "css", "html"}, 10, true)
	ontology := createOntology([]string{"js", "css", "html"}, false, false)

	errs := checkRawData(ontology, raw)

	err := checkNormalizerErrorCode(errs, normalizerErrors.CaseCodeOntologyMissingRuleForSameSkill)
	if err != nil {
		t.Error(err)
	}
}

func TestResolveDublicates_HasDublicates_ReturnsNoDublicates(t *testing.T) {
	raw := createRawData([]string{"js", "css", "html"}, 10, false)
	plainData := getPlainData(raw)

	stats, _ := resolveDublicates(plainData)

	idsMap := make(map[string]bool)

	for _, stat := range stats {
		if _, ok := idsMap[stat.ID]; ok {
			t.Error("Collection has dublicated ID: " + stat.ID)
			t.FailNow()
		} else {
			idsMap[stat.ID] = true
		}
	}
}

func TestResolveDublicates_HasDublicates_TotalCountEqualsActualCount(t *testing.T) {
	raw := createRawData([]string{"js", "css", "html"}, 10, false)
	plainData := getPlainData(raw)

	_, totalCount := resolveDublicates(plainData)
	actualDublicates := 0
	idsMap := make(map[string]bool)

	for _, stat := range plainData {
		if _, ok := idsMap[stat.ID]; ok {
			actualDublicates++
		} else {
			idsMap[stat.ID] = true
		}
	}

	if actualDublicates != totalCount {
		t.Errorf("Expect %d dublicates, but have %d", totalCount, actualDublicates)
	}
}

func TestNormalizeRawData_IsCorrect_ReturnsCollection(t *testing.T) {
	raw := createRawData([]string{"js", "css", "html"}, 10, true)
	ontology := createOntology([]string{"js", "css", "html"}, false, true)

	data, _ := NormalizeRawData(ontology, raw)

	if len(data) == 0 {
		t.Error("Empty collection")
	}
}

func TestNormalizeRawData_IsNotCorrect_ReturnsError(t *testing.T) {
	raw := createRawData([]string{"js", "css", "html"}, 10, true)
	ontology := createOntology([]string{"js", "css"}, true, false)

	_, errs := NormalizeRawData(ontology, raw)

	if len(errs) == 0 {
		t.Errorf("Has no errors")
	}
}
