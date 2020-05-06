package userconfig

import "time"

type activity struct {
	name        string
	periodicity periodicity
	grow        grow
}

func (act activity) IsActual(startTime time.Time, currentTime time.Time) bool {
	passedInMetric := act.passedInMetric(startTime, currentTime)
	result := (passedInMetric + act.periodicity.denominator) % act.periodicity.addendum
	return !(result == 0)
}

func (act activity) GetSummary(startTime time.Time, currentTime time.Time) string {
	passedInMetric := act.passedInMetric(startTime, currentTime)
	grow := act.getGrow(passedInMetric)
	return string(grow) + " " + act.name
}

func (act activity) getGrow(passedInMetric int) int {
	switch act.grow.growFunction._type {
	case "monotonous":
		return act.getMonotonousGrow(passedInMetric)
	default:
		return act.getConstGrow(passedInMetric)
	}
}

func (act activity) getConstGrow(passedInMetric int) int {
	return act.grow.borders.leftBorder
}

func (act activity) getMonotonousGrow(passedInMetric int) int {
	expected := act.grow.borders.leftBorder + act.grow.growFunction.coefficient*passedInMetric
	if expected > act.grow.borders.rightBorder {
		return act.grow.borders.rightBorder
	}
	return expected
}

func (act activity) passedInMetric(startTime time.Time, currentTime time.Time) int {
	passedInMetric := int(currentTime.Sub(startTime).Hours())
	if act.periodicity.metrics == "day" {
		passedInMetric = int(passedInMetric / 24)
	}
	return passedInMetric
}

type grow struct {
	borders      borders
	growFunction growFunction
}

type growFunction struct {
	_type       string
	coefficient int
}

type borders struct {
	_type       string
	leftBorder  int
	rightBorder int
}

type periodicity struct {
	metrics     string
	denominator int
	addendum    int
}
