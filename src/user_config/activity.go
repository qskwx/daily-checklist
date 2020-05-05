package userconfig

import "time"

type activity struct {
	Name        string
	Periodicity periodicity
	Grow        grow
}

func (act activity) IsActual(startTime time.Time, currentTime time.Time) bool {
	passedInMetric := int(currentTime.Sub(startTime).Hours())
	if act.Periodicity.Metrics == "day" {
		passedInMetric = int(passedInMetric / 24)
	}
	result := (passedInMetric + act.Periodicity.Denominator) % act.Periodicity.Addendum
	return !(result == 0)
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
