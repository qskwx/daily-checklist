package session

import (
	"fmt"
	"strings"
)

type categories []category

func categoriesFabric(activities map[string][]string) []category {
	categories := make([]category, 0)
	categoryIdx := 0
	for name, acts := range activities {
		prefix := toCharStr(categoryIdx)
		categories = append(categories, categoryFabric(name, prefix, acts))
		categoryIdx++
	}
	return categories
}

func (cats *categories) setDone(mixedID string) error {
	for _, category := range *cats {
		if strings.HasPrefix(mixedID, category.prefix) {
			return category.setDone(mixedID)
		}
	}
	return fmt.Errorf("unable to find category for activity id = '%s'", mixedID)
}

func toCharStr(i int) string {
	return string('a' + i)
}
