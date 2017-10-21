package cleaner

import (
	"testing"

	"github.com/AxelUser/gowork/errors/normalizerErrors"
	"github.com/AxelUser/gowork/utils"
)

func TestCheckRawData_NoData_ReturnsError(t *testing.T) {
	raw := utils.CreateRawData([]string{"js", "css"}, 10, true)
	ontology := utils.CreateOntology([]string{"js", "css", "html"}, false, true)

	errs := checkRawData(ontology, raw)

	err := utils.CheckNormalizerErrorCode(errs, normalizerErrors.CaseCodeMissingData)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckRawData_EmptyRules_ReturnsErrors(t *testing.T) {
	raw := utils.CreateRawData([]string{"js", "css", "html"}, 10, true)
	ontology := utils.CreateOntology([]string{"js", "css", "html"}, true, false)

	errs := checkRawData(ontology, raw)

	err := utils.CheckNormalizerErrorCode(errs, normalizerErrors.CaseCodeEmptyRules)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckRawData_MissingRulesForSameSkill_ReturnsError(t *testing.T) {
	raw := utils.CreateRawData([]string{"js", "css", "html"}, 10, true)
	ontology := utils.CreateOntology([]string{"js", "css", "html"}, false, false)

	errs := checkRawData(ontology, raw)

	err := utils.CheckNormalizerErrorCode(errs, normalizerErrors.CaseCodeOntologyMissingRuleForSameSkill)
	if err != nil {
		t.Error(err)
	}
}

func TestResolveDublicates_HasDublicates_ReturnsNoDublicates(t *testing.T) {
	raw := utils.CreateRawData([]string{"js", "css", "html"}, 10, false)
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
	raw := utils.CreateRawData([]string{"js", "css", "html"}, 10, false)
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

func TestResolveDublicates_HasDublicates_AddsSkillsToDublicates(t *testing.T) {
	raw := utils.CreateRawData([]string{"js", "css", "html"}, 1, false)
	plainData := getPlainData(raw)

	dataWithoutDublicates, _ := resolveDublicates(plainData)

	for _, stat := range dataWithoutDublicates {
		if len(stat.Skills) != 3 {
			t.Errorf("Expect 3 skills for ID <%s>, but have %d", stat.ID, len(stat.Skills))
		}
	}
}
