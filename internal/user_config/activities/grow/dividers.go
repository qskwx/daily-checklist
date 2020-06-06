package grow

import (
	"fmt"
)

type dividers []int

func (div dividers) format() string {
	return "%v"
}

func (div dividers) items(amount int) []int {
	items := make([]int, len(div))
	copy(items, div)
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
