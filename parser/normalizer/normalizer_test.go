package normalizer

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/AxelUser/gowork/errors/normalizerErrors"
	"github.com/AxelUser/gowork/models"
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

func checkNormalizerErrorCode(errs []error, code int) error {
	if len(errs) == 1 {
		switch errs[0].(type) {
		case normalizerErrors.NormalizerError:
			e := errs[0].(normalizerErrors.NormalizerError)
			if e.CaseCode == code {
				return nil
			}
			return fmt.Errorf("NormalizerError is must be with code %d: %d", code, e.CaseCode)
		default:
			return fmt.Errorf("Error is not NormalizerError: %s", reflect.TypeOf(errs[0]))
		}
	} else {
		return fmt.Errorf("There are multiple errors: %d", len(errs))
	}
}

func TestNormalizeRawDataNoData(t *testing.T) {
	raw := createRawData([]string{"js", "css"}, 10)
	ontology := createOntology([]string{"js", "css", "html"}, false)

	_, errs := NormalizeRawData(ontology, raw)

	err := checkNormalizerErrorCode(errs, normalizerErrors.CaseCodeMissingData)
	if err != nil {
		t.Error(err)
	}
}

func TestNormalizeRawDataEmptyRules(t *testing.T) {
	raw := createRawData([]string{"js", "css", "html"}, 10)
	ontology := createOntology([]string{"js", "css", "html"}, true)

	_, errs := NormalizeRawData(ontology, raw)

	err := checkNormalizerErrorCode(errs, normalizerErrors.CaseCodeEmptyRules)
	if err != nil {
		t.Error(err)
	}
}

func TestNormalizeRawDataHasDublicates(t *testing.T) {
	t.Error("Not implemented")
}

func TestNormalizeRawDataMissingRulesForSameSkill(t *testing.T) {
	t.Error("Not implemented")
}

func TestNormalizeRawData(t *testing.T) {
	raw := createRawData([]string{"js", "css", "html"}, 10)
	ontology := createOntology([]string{"js", "css", "html"}, false)

	data, _ := NormalizeRawData(ontology, raw)

	if data == nil {
		t.Error("Empty collection")
	}
}
