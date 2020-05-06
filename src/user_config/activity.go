package userconfig

import "time"

type activity struct {
	Name        string
	Periodicity periodicity
	Grow        grow
}

func (act activity) IsActual(startTime time.Time, currentTime time.Time) bool {
	passedInMetric := act.passedInMetric(startTime, currentTime)
	result := (passedInMetric + act.Periodicity.Denominator) % act.Periodicity.Addendum
	return !(result == 0)
}

func (act activity) GetSummary(startTime time.Time, currentTime time.Time) bool {
	passedInMetric := act.passedInMetric(startTime, currentTime)
	result := (passedInMetric + act.Periodicity.Denominator) % act.Periodicity.Addendum
	return !(result == 0)
}

func (act activity) getGrow(passedInMetric int) int {
	switch act.Grow.GrowFunction.Type {
	case "monotonous":
		return act.getMonotonousGrow(passedInMetric)
	default:
		return act.getConstGrow(passedInMetric)
	}
}

func (act activity) getConstGrow(passedInMetric int) int {
	return act.Grow.Borders.LeftBorder
}

func (act activity) getMonotonousGrow(passedInMetric int) int {
	expected := act.Grow.Borders.LeftBorder + act.Grow.GrowFunction.Coefficient*passedInMetric
	if expected > act.Grow.Borders.RightBorder {
		return act.Grow.Borders.RightBorder
	}
	return expected
}

func (act activity) passedInMetric(startTime time.Time, currentTime time.Time) int {
	passedInMetric := int(currentTime.Sub(startTime).Hours())
	if act.Periodicity.Metrics == "day" {
		passedInMetric = int(passedInMetric / 24)
	}
	return passedInMetric
}

type grow struct {
	Borders      borders
	GrowFunction growFunction
}

type growFunction struct {
	Type        string
	Coefficient int
}

type borders struct {
	Type        string
	LeftBorder  int
	RightBorder int
}

type periodicity struct {
	Metrics     string
	Denominator int
	Addendum    int
}
