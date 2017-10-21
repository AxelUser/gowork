package normalizer

import (
	"github.com/AxelUser/gowork/models/configs"
	"github.com/AxelUser/gowork/models/dataModels"
	"github.com/AxelUser/gowork/parser/normalizer/cleaner"
	"github.com/AxelUser/gowork/parser/normalizer/selector"
)

func normalizeInputsAndOutputs(ontology []configs.OntologyData, data []dataModels.VacancyStats) ([]float32, []float32) {
	return nil, nil
}

// NormalizeRawData proceeds vacancies and normalize them for training set
func NormalizeRawData(ontologyInfos []configs.OntologyData, rawData map[string][]dataModels.VacancyStats) ([]dataModels.TraingingSetItem, []error) {
	uniqueData, errs := cleaner.Clean(ontologyInfos, rawData)
	if len(errs) > 0 {
		return nil, errs
	}

	return selector.SelectSkills(ontologyInfos, uniqueData), nil
}
