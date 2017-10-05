package normalizerErrors

import (
	"fmt"
	"strings"
)

// CaseCodeMissingData is for error when ontology have skills with no data for them
const CaseCodeMissingData int = 0

// CaseCodeEmptyRules is for error when ontology for skill has empty rules
const CaseCodeEmptyRules int = 1

// CaseCodeOntologyMissingRuleForSameSkill is for error when ontology for skill A missing rule for skill A
const CaseCodeOntologyMissingRuleForSameSkill = 2

// NormalizerError is for errors during normalizing data
type NormalizerError struct {
	CaseCode         int
	AliasesWithError []string
	InnerError       error
}

func (e NormalizerError) Error() string {
	aliases := strings.Join(e.AliasesWithError, ", ")
	switch e.CaseCode {
	case CaseCodeMissingData:
		return "Missing for skills in ontology: " + aliases
	case CaseCodeEmptyRules:
		return "Empty rules in ontology: " + aliases
	case CaseCodeOntologyMissingRuleForSameSkill:
		return "Missing rules for skills themselves: " + aliases
	default:
		return fmt.Sprintf("Unexpected error: %s", e.InnerError)
	}
}

// New creates NormalizerError
func New(code int, aliases []string, inner error) NormalizerError {
	return NormalizerError{CaseCode: code, AliasesWithError: aliases, InnerError: inner}
}
