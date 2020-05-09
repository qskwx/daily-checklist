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
		// probably index error
		if strings.HasPrefix(mixedID, category.prefix) {
			activityID := mixedID[len(category.prefix):]
			if err := cats.setDone(activityID); err != nil {
				return err
			}
		}
	}
	return fmt.Errorf("unable to find category for activity id = '%s'", mixedID)
}
