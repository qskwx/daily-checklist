package sections

import (
	"daily-checklist/internal/user_config/activities"
	"time"
)

type section struct {
	Name       string
	Activities activities.Activities
}

func (sec section) GetCurrentActivities(startTime time.Time, currentTime time.Time) []string {
	sectionActivities := make([]string, 0)
	for _, act := range sec.Activities {
		if act.IsActual(startTime, currentTime) {
			sectionActivities = append(sectionActivities, act.GetSummary(startTime, currentTime))
		}
	}
	return sectionActivities
}
