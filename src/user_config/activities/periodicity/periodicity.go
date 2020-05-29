package periodicity

type Periodicity struct {
	denominator int
	addendum    int
}

func (period Periodicity) IsActual(passedInMetric int) bool {
	result := (passedInMetric + period.addendum) % period.denominator
	return result == 0
}
