package session

import (
	"fmt"
	"strconv"
)

type activities []activity

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

func (acts *activities) setDone(actID string) error {
	for idx := range *acts {
		if (*acts)[idx].Id() == actID {
			return (*acts)[idx].SetDone()
		}
	}
	return fmt.Errorf("unable to found activity with id = '%s'", actID)
}
