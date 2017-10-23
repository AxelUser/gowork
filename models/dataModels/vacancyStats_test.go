package dataModels

import "testing"

func TestNewVacancyStats_PassNullToBothSalaries_SalariesEqualZero(t *testing.T) {
	stat := NewVacancyStats("1", "example.com", nil, nil, "RUB", "css")
	if stat.SalaryFrom != 0 || stat.SalaryTo != 0 {
		t.Error("Salaries expected to be zeros")
	}
}

func TestNewVacancyStats_PassNotNullToBothSalaries_InitSalaries(t *testing.T) {
	salaryFrom := 1000.0
	salaryTo := 10000.0
	stat := NewVacancyStats("1", "example.com", &salaryFrom, &salaryTo, "RUB", "css")
	if stat.SalaryFrom != salaryFrom || stat.SalaryTo != salaryTo {
		t.Errorf("Salaries expected to be %f and %f, but were %f and %f",
			salaryFrom, salaryTo, stat.SalaryFrom, salaryTo)
	}
}

func TestNewVacancyStats_PassNotNullToSalaryFrom_SalaryFromMustBeEqualSalaryTo(t *testing.T) {
	salaryTo := 10000.0
	stat := NewVacancyStats("1", "example.com", nil, &salaryTo, "RUB", "css")
	if stat.SalaryFrom != stat.SalaryTo {
		t.Errorf("SalaryFrom expected to be %f, but was %f", stat.SalaryTo, stat.SalaryFrom)
	}
}

func TestNewVacancyStats_PassNotNullToSalaryTo_SalaryToMustBeEqualSalaryFrom(t *testing.T) {
	salaryFrom := 10000.0
	stat := NewVacancyStats("1", "example.com", &salaryFrom, nil, "RUB", "css")
	if stat.SalaryTo != stat.SalaryFrom {
		t.Errorf("SalaryTo expected to be %f, but was %f", stat.SalaryFrom, stat.SalaryTo)
	}
}
