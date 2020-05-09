package session

import "strconv"

type category struct {
	name       string
	prefix     string
	activities []activity
}

func categoryFabric(name string, prefix string, acts []string) category {
	return category{
		name:       name,
		prefix:     prefix,
		activities: activitiesFabric(prefix, acts),
	}
}

func activitiesFabric(prefix string, acts []string) []activity {
	activities := make([]activity, 0, 8)
	for idx, act := range acts {
		activities = append(activities, activityFabric(
			prefix+strconv.Itoa(idx),
			act,
			false,
		))
	}
	return activities
}
