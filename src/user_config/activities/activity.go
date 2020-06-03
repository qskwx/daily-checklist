package activities

import (
	"daily-checklist/src/user_config/activities/grow"
	"daily-checklist/src/user_config/activities/periodicity"
	"time"
)

type Activity struct {
	Name        string
	Periodicity periodicity.Periodicity
	Grow        grow.Grow
}

func (act Activity) IsActual(startTime time.Time, currentTime time.Time) bool {
	passedInMetric := act.passedInMetric(startTime, currentTime)
	return act.Periodicity.IsActual(passedInMetric)
}

func (act Activity) GetSummary(startTime time.Time, currentTime time.Time) string {
	passedInMetric := act.passedInMetric(startTime, currentTime)
	grow := act.Grow.Show(passedInMetric)
	summary := ""
	if grow != "" {
		summary += grow + " "
	}
	summary += act.Name
	return summary
}

func (act Activity) passedInMetric(startTime time.Time, currentTime time.Time) int {
	passedInMetric := int(currentTime.Sub(startTime).Hours() / 24)
	return passedInMetric
}
