package normalizer

import (
	"testing"

	"github.com/AxelUser/gowork/utils"
)

func TestNormalizeRawData_IsCorrect_ReturnsCollection(t *testing.T) {
	raw := utils.CreateRawData([]string{"js", "css", "html"}, 10, true)
	ontology := utils.CreateOntology([]string{"js", "css", "html"}, false, true)

	data, _ := NormalizeRawData(ontology, raw)

	if len(data) == 0 {
		t.Error("Empty collection")
	}
}

func TestNormalizeRawData_IsNotCorrect_ReturnsError(t *testing.T) {
	raw := utils.CreateRawData([]string{"js", "css", "html"}, 10, true)
	ontology := utils.CreateOntology([]string{"js", "css"}, true, false)

	_, errs := NormalizeRawData(ontology, raw)

	if len(errs) == 0 {
		t.Errorf("Has no errors")
	}
}
