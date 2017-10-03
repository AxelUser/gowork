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
	}
	return fmt.Sprintf("Error. %s", ev.Error)
}

// IsSuccess returns true, if there is no error
func (ev DataLoadedEvent) IsSuccess() bool {
	return ev.Error == nil
}

// NewDataLoadedEvent creates new DataLoadedEvent
func NewDataLoadedEvent(skill string, url string, data []models.VacancyStats) DataLoadedEvent {
	return DataLoadedEvent{Skill: skill, URL: url, Data: data, Error: nil}
}

// NewDataLoadedEventWithError creates new DataLoadedEvent with error object
func NewDataLoadedEventWithError(skill string, url string, err error) DataLoadedEvent {
	return DataLoadedEvent{Skill: skill, URL: url, Data: nil, Error: err}
}
