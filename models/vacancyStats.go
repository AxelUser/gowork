package models

// VacancyStats is for data about salary for vacancy
type VacancyStats struct {
	ID         string
	URL        string
	SalaryFrom *int
	SalaryTo   *int
	Currency   string
	Skills     []string
}

// AddSkill adds skill to collections
func (s VacancyStats) AddSkill(skill string) {
	s.Skills = append(s.Skills, skill)
}

// NewVacancyStats creates VacancyStats
func NewVacancyStats(id string, url string, salaryFrom *int, salaryTo *int, currency string, skills ...string) VacancyStats {
	return VacancyStats{
		ID:         id,
		URL:        url,
		SalaryFrom: salaryFrom,
		SalaryTo:   salaryTo,
		Currency:   currency,
		Skills:     skills}
}
