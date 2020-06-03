package grow

type Grow struct {
	Borders      borders
	GrowFunction growFunction `json:"grow-function,"`
	ShowF        show
}

func (grow Grow) Show(passedInMetric int) string {
	growCount := grow.count(passedInMetric)
	return grow.ShowF.show(growCount)
}

func (grow Grow) count(passedInMetric int) int {
	switch grow.GrowFunction.GrowType {
	case "monotonous":
		return grow.getMonotonousGrow(passedInMetric)
	default:
		return grow.getConstGrow(passedInMetric)
	}
}

func (grow Grow) getConstGrow(passedInMetric int) int {
	return grow.Borders.Left
}

func (grow Grow) getMonotonousGrow(passedInMetric int) int {
	expected := grow.Borders.Left + int(grow.GrowFunction.Coefficient*float64(passedInMetric))
	if expected > grow.Borders.Right {
		return grow.Borders.Right
	}
	return expected
}
