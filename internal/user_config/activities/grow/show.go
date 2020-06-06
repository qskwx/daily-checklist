package grow

import (
	"fmt"
)

type show struct {
	Dividers dividers
}

func (show show) show(amount int) string {
	summary := ""
	if amount != 0 {
		format := "%d"
		if show.Dividers != nil {
			format += " "
			format += show.Dividers.show(amount)
		}
		summary += fmt.Sprintf(format, amount)
	}
	return summary
}
