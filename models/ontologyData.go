package models

// OntologyData stores rules and main info for loaded skill
type OntologyData struct {
	Alias   string `json:"alias"`
	Caption string `json:"caption"`
	Rules   struct {
		Javascript float64 `json:"javascript"`
	} `json:"rules"`
}
