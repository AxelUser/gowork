package events

import (
	"fmt"

	"github.com/AxelUser/gowork/models"
)

// DataLoadedEvent is event for loaded data from HeadHunter API
type DataLoadedEvent struct {
	Skill string
	URL   string
	Data  []models.VacancyStats
	Error error
}

func (ev DataLoadedEvent) String() string {
	if ev.IsSuccess() {
		return fmt.Sprintf("Loaded data (%d item(s)) for skill '%s' from %s", len(ev.Data), ev.Skill, ev.URL)
	} else if ev.HasData() {
		return fmt.Sprintf("Loaded data (%d item(s)) for skill '%s' from %s. Error: %s", len(ev.Data), ev.Skill, ev.URL, ev.Error)
	}
	return fmt.Sprintf("Error. %s", ev.Error)
}

// IsSuccess returns true, if there is no error
func (ev DataLoadedEvent) IsSuccess() bool {
	return ev.Error == nil
}

// HasData returns true, if some data was loaded
func (ev DataLoadedEvent) HasData() bool {
	return len(ev.Data) > 0
}

// NewDataLoadedEvent creates new DataLoadedEvent
func NewDataLoadedEvent(skill string, url string, data []models.VacancyStats) DataLoadedEvent {
	return DataLoadedEvent{Skill: skill, URL: url, Data: data, Error: nil}
}

// NewDataLoadedEventWithError creates new DataLoadedEvent with error object
func NewDataLoadedEventWithError(skill string, url string, data []models.VacancyStats, err error) DataLoadedEvent {
	return DataLoadedEvent{Skill: skill, URL: url, Data: data, Error: err}
}
