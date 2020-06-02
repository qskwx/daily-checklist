package periodicity

type Periodicity struct {
	Denominator int
	Addendum    int
}

func (period Periodicity) IsActual(passedInMetric int) bool {
	result := (passedInMetric + period.Addendum) % period.Denominator
	return result == 0
}
