package userconfig

import (
	"strconv"
	"time"
)

type Activity struct {
	Name        string
	Periodicity Periodicity
	Grow        Grow
}

func (act Activity) IsActual(startTime time.Time, currentTime time.Time) bool {
	passedInMetric := act.passedInMetric(startTime, currentTime)
	result := (passedInMetric + act.Periodicity.Addendum) % act.Periodicity.Denominator
	return result == 0
}

func (act Activity) GetSummary(startTime time.Time, currentTime time.Time) string {
	passedInMetric := act.passedInMetric(startTime, currentTime)
	grow := act.getGrow(passedInMetric)
	summary := ""
	if grow != 0 {
		summary += strconv.Itoa(grow) + " "
	}
	summary += act.Name
	return summary
}

func (act Activity) getGrow(passedInMetric int) int {
	switch act.Grow.GrowFunction.Type {
	case "monotonous":
		return act.getMonotonousGrow(passedInMetric)
	default:
		return act.getConstGrow(passedInMetric)
	}
}

func (act Activity) getConstGrow(passedInMetric int) int {
	return act.Grow.Borders.Left
}

func (act Activity) getMonotonousGrow(passedInMetric int) int {
	expected := act.Grow.Borders.Left + act.Grow.GrowFunction.Coefficient*passedInMetric
	if expected > act.Grow.Borders.Right {
		return act.Grow.Borders.Right
	}
	return expected
}

func (act Activity) passedInMetric(startTime time.Time, currentTime time.Time) int {
	passedInMetric := int(currentTime.Sub(startTime).Hours() / 24)
	return passedInMetric
}

type Grow struct {
	Borders      Borders
	GrowFunction GrowFunction `json:"grow-function,"`
}

type GrowFunction struct {
	Type        string
	Coefficient int
}

type Borders struct {
	Left  int
	Right int
}

type Periodicity struct {
	Denominator int
	Addendum    int
}
