package userconfig

import "time"

type activity struct {
	Name        string
	Periodicity periodicity
	Grow        grow
}

func (act activity) IsActual(startDate time.Time, currentDate time.Time) bool {
	passedInMetric := int(currentDate.Sub(startDate).Hours())
	if act.Periodicity.Metrics == "day" {
		passedInMetric = int(passedInMetric / 24)
	}
	denominator := act.Periodicity.Denominator
	addendum := act.Periodicity.Addendum
	result := (passedInMetric + addendum) % denominator
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
