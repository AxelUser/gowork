package events

import (
	"fmt"

	"gowork/models"
)

// DataLoadedEvent is event for loaded data from HeadHunter API
type DataLoadedEvent struct {
	Skill string
	URL   string
	Data  []models.VacancyStats
}

func (ev DataLoadedEvent) String() string {
	return fmt.Sprintf("Loaded data (%d item(s)) for skill '%s' from %s", len(ev.Data), ev.Skill, ev.URL)
}

// NewDataLoadedEvent creates new DataLoadedEvent
func NewDataLoadedEvent(skill string, url string, data []models.VacancyStats) DataLoadedEvent {
	return DataLoadedEvent{Skill: skill, URL: url, Data: data}
}
