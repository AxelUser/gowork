package normalizer_tests

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/AxelUser/gowork/errors/normalizerErrors"
	"github.com/AxelUser/gowork/models"
	"github.com/AxelUser/gowork/parser/normalizer"
)

func createRawData(aliases []string, countPerSkill int) map[string][]models.VacancyStats {
	statsMap := make(map[string][]models.VacancyStats)
	for i, alias := range aliases {
		var stats []models.VacancyStats
		for j := 0; i < countPerSkill; i++ {
			salaryFrom := (i + 1) * 10000
			salaryTo := (i + 1) * 20000
			s := models.NewVacancyStats(strconv.Itoa((i+1)*100000+j), "test.com", &salaryFrom, &salaryTo, "RUB")
			stats = append(stats, s)
		}
		statsMap[alias] = stats
	}
	return statsMap
}

func createOntology(aliases []string, emptyRules bool) []models.OntologyData {
	var ontology []models.OntologyData
	for _, alias := range aliases {
		o := models.OntologyData{Alias: alias, Caption: alias}
		if !emptyRules {
			o.Rules = make(map[string]float32)
			o.Rules[alias] = 1
		}
		ontology = append(ontology, o)
	}

	return ontology
}

// TestCheckRawDataNoData is test for checkRawData
func TestCheckRawDataNoData(t *testing.T) {
	raw := createRawData([]string{"js", "css"}, 10)
	ontology := createOntology([]string{"js", "css", "html"}, false)

	_, errs := normalizer.NormalizeRawData(ontology, raw)

	if len(errs) == 1 {
		switch errs[0].(type) {
		case normalizerErrors.NormalizerError:
			e := errs[0].(normalizerErrors.NormalizerError)
			if e.CaseCode == normalizerErrors.CaseCodeMissingData {
				return
			}
			t.Errorf("NormalizerError is must be with code %d: %d", normalizerErrors.CaseCodeMissingData, e.CaseCode)
		default:
			t.Errorf("Error is not NormalizerError: %s", reflect.TypeOf(errs[0]))
		}
	} else {
		t.Errorf("There are multiple errors: %d", len(errs))
	}
}
