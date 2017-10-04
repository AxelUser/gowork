package models

// OntologyData stores rules and main info for loaded skill
type OntologyData struct {
	Alias   string             `json:"alias"`
	Caption string             `json:"caption"`
	Rules   map[string]float32 `json:"rules"`
}
