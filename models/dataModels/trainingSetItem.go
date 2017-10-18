package dataModels

// TraingingSetItem contains skills for inputs and salary for output
type TraingingSetItem struct {
	SkillsLevels []float32
	SalaryFrom   float32
	SalaryTo     float32
}

// NewTraingingSetItem creates new TraingingSetItem
func NewTraingingSetItem(from float32, to float32, skills []float32) TraingingSetItem {
	return TraingingSetItem{SalaryFrom: from, SalaryTo: to, SkillsLevels: skills}
}
