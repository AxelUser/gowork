package models

// ParserConfig represents config for parser
type ParserConfig struct {
	URL      string            `json:"url"`
	Defaults map[string]string `json:"defaults"`
	Queries  []ParserQuery     `json:"queries"`
}
