package normalizer

import (
	"github.com/AxelUser/gowork/models/configs"
	"github.com/AxelUser/gowork/models/dataModels"
	"github.com/AxelUser/gowork/parser/normalizer/cleaner"
	"github.com/AxelUser/gowork/parser/normalizer/selector"
)

const maxVacancySalaryInRubles = 1000000

// NormalizeRawData proceeds vacancies and normalize them for training set
func NormalizeRawData(ontologyInfos []configs.OntologyData, rawData map[string][]dataModels.VacancyStats) ([]dataModels.TraingingSetItem, []error) {
	uniqueData, errs := cleaner.Clean(ontologyInfos, rawData)
	if len(errs) > 0 {
		return nil, errs
	}

	return selector.SelectSkills(ontologyInfos, uniqueData), nil
}
