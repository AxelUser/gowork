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

func TestGetVectorForSkills_PassCorrectParameters_ReturnsVector(t *testing.T) {
	skillsOrder := []string{"js", "css", "html"}
	rawData := utils.CreateRawData([]string{"js"}, 1, true)
	ontology := utils.CreateOntology(skillsOrder, false, true)
	ontology[0].Rules["css"] = 0.3

	vector := getVectorForSkills(ontology, skillsOrder, rawData["js"][0].Skills)

	jsRule := ontology[0].Rules["js"]
	cssRule := ontology[0].Rules["css"]
	htmlRule := ontology[0].Rules["html"]

	if vector[0] != jsRule || vector[1] != cssRule || vector[2] != htmlRule {
		t.Errorf("Expected vector [%f, %f, %f], but was [%f, %f, %f]",
			jsRule, cssRule, htmlRule, vector[0], vector[1], vector[2])
	}
}
