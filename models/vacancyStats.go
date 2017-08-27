package models

// VacancyStats is for data about salary for vacancy
type VacancyStats struct {
	ID           string
	URL          string
	Technologies map[string]string
	SalaryFrom   *int
	SalaryTo     *int
	Currency     string
}

// NewVacancyStats creates VacancyStats
func NewVacancyStats(id string, url string, salaryFrom *int, salaryTo *int, currency string) VacancyStats {
	return VacancyStats{
		ID:           id,
		URL:          url,
		Technologies: make(map[string]string),
		SalaryFrom:   salaryFrom,
		SalaryTo:     salaryTo}
}
