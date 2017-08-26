package models

// VacancyStats is for data about salary for vacancy
type VacancyStats struct {
	ID           string
	URL          string
	Technologies map[string]string
	SalaryFrom   *float32
	SalaryTo     *float32
}

// NewVacancyStats creates VacancyStats
func NewVacancyStats(id string, url string, salaryFrom *float32, salaryTo *float32) VacancyStats {
	return VacancyStats{
		ID:           id,
		URL:          url,
		Technologies: make(map[string]string),
		SalaryFrom:   salaryFrom,
		SalaryTo:     salaryTo}
}
