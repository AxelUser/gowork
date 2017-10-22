package api

// CurrencyRates stores latest currency rates for base currency
type CurrencyRates struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float32 `json:"rates"`
	Error string             `json:"error"`
}
