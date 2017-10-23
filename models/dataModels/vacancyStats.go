package dataModels

// VacancyStats is for data about salary for vacancy
type VacancyStats struct {
	ID         string
	URL        string
	SalaryFrom float64
	SalaryTo   float64
	Currency   string
	Skills     []string
}

// AddSkill adds skill to collections
func (s *VacancyStats) AddSkill(skill string) {
	s.Skills = append(s.Skills, skill)
}

// NewVacancyStats creates VacancyStats
func NewVacancyStats(id string, url string, salaryFrom *float64, salaryTo *float64, currency string, skills ...string) VacancyStats {
	var restoredSalaryFrom float64
	var restoredSalaryTo float64

	if salaryFrom == nil && salaryTo != nil {
		restoredSalaryFrom = *salaryTo
	} else if salaryFrom != nil {
		restoredSalaryFrom = *salaryFrom
	}

	if salaryTo == nil && salaryFrom != nil {
		restoredSalaryTo = *salaryFrom
	} else if salaryTo != nil {
		restoredSalaryTo = *salaryTo
	}

	return VacancyStats{
		ID:         id,
		URL:        url,
		SalaryFrom: restoredSalaryFrom,
		SalaryTo:   restoredSalaryTo,
		Currency:   currency,
		Skills:     skills}
}
