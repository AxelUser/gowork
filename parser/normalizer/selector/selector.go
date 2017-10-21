package selector

import (
	"github.com/AxelUser/gowork/models/configs"
	"github.com/AxelUser/gowork/models/dataModels"
)

const maxVectorValue = 1

// SelectSkills selects skills for vacancies and returns training set with normalized skills vector
func SelectSkills(ontology []configs.OntologyData, stats []dataModels.VacancyStats) []dataModels.TraingingSetItem {
	var trainingSet []dataModels.TraingingSetItem
	skillsOrder := getSkillsOrder(ontology)

	for _, stat := range stats {
		vSkills := getVectorForSkills(ontology, skillsOrder, stat.Skills)
		trainingSet = append(trainingSet, dataModels.NewTraingingSetItem(0, 0, vSkills))
	}

	return trainingSet
}

func getVectorForSkills(ontology []configs.OntologyData, skillsOrder []string, skills []string) []float32 {
	vector := make([]float32, len(ontology))

	skillsHashSet := make(map[string]bool)

	for _, skill := range skills {
		skillsHashSet[skill] = true
	}

	for _, oSet := range ontology {
		if _, ok := skillsHashSet[oSet.Alias]; ok {
			for i, skillInOrder := range skillsOrder {
				ruleValue := oSet.Rules[skillInOrder]
				newValue := vector[i] + ruleValue
				if newValue <= maxVectorValue {
					vector[i] = newValue
				}
			}
		}
	}

	return vector
}

func getSkillsOrder(onotlogy []configs.OntologyData) []string {
	skillsOrder := make([]string, len(onotlogy))
	for i, oSet := range onotlogy {
		skillsOrder[i] = oSet.Alias
	}
	return skillsOrder
}
