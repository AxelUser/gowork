package selector

import "testing"
import "github.com/AxelUser/gowork/utils"

func TestGetSkillsOrder_PassOntology_SkillsAreInSameOrder(t *testing.T) {
	ontology := utils.CreateOntology([]string{"js", "css", "html"}, false, true)

	skillsInOrder := getSkillsOrder(ontology)

	for i := range ontology {
		if ontology[i].Alias != skillsInOrder[i] {
			t.Errorf("Expected skill <%s>, but was <%s>", ontology[i].Alias, skillsInOrder[i])
		}
	}
}

func TestGetSkillsOrder_PassOntology_SameLenght(t *testing.T) {
	ontology := utils.CreateOntology([]string{"js", "css", "html"}, false, true)

	skillsInOrder := getSkillsOrder(ontology)

	if len(ontology) != len(skillsInOrder) {
		t.Errorf("Expected length <%d>, but was <%d>", len(ontology), len(skillsInOrder))
	}
}
