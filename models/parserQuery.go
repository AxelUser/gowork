package models

// ParserQuery represents query for api
type ParserQuery struct {
	Alias string `json:"alias"`
	Text  string `json:"query_text"`
}
