package models

// ParserConfig represents config for parser
type ParserConfig struct {
	Host     string            `json:"host"`
	Pathname string            `json:"pathname"`
	Defaults map[string]string `json:"defaults"`
	Queries  []ParserQuery     `json:"queries"`
}

// ParserQuery represents query for api
type ParserQuery struct {
	Alias string `json:"alias"`
	Text  string `json:"text"`
}
