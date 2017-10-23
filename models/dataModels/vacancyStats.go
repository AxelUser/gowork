package dataModels

// VacancyStats is for data about salary for vacancy
type VacancyStats struct {
	ID         string
	URL        string
	SalaryFrom float32
	SalaryTo   float32
	Currency   string
	Skills     []string
}

// AddSkill adds skill to collections
func (s *VacancyStats) AddSkill(skill string) {
	s.Skills = append(s.Skills, skill)
}

// NewVacancyStats creates VacancyStats
func NewVacancyStats(id string, url string, salaryFrom *float32, salaryTo *float32, currency string, skills ...string) VacancyStats {
	var restoredSalaryFrom float32
	var restoredSalaryTo float32

	if salaryFrom == nil && salaryTo != nil {
		restoredSalaryFrom = *salaryTo
	} else {
		restoredSalaryFrom = *salaryFrom
	}

	if salaryTo == nil && salaryFrom != nil {
		restoredSalaryTo = *salaryFrom
	} else {
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
