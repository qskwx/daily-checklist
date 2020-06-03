package grow

import (
	"fmt"
	"strings"
)

type dividers []int

func (div dividers) format() string {
	ln := make([]string, 0)
	for range div {
		ln = append(ln, "%d")
	}
	return strings.Join(ln, ",")
}

func (div dividers) items(amount int) []int {
	var items []int
	copy(div, items)
	for idx := range items {
		items[idx] = int((float32(amount)/float32(div[idx]) + 0.5))
	}
	return items
}

func (div dividers) show(amount int) string {
	ft := div.format()
	items := div.items(amount)
	return fmt.Sprintf(ft, items)
}
