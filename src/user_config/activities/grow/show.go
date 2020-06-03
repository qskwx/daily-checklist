package grow

import (
	"fmt"
)

type show struct {
	Dividers dividers
}

func (show show) show(amount int) string {
	format := "%d"
	if show.Dividers != nil {
		format += " ["
		format += show.Dividers.show(amount)
		format += "]"
	}
	return fmt.Sprintf(format, amount)
}
