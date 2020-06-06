package sections

import "time"

type Sections []section

func (sec Sections) GetCurrentActivities(startTime time.Time, currentTime time.Time) map[string][]string {
	currentActivities := make(map[string][]string)
	for _, section := range sec {
		sectionActivities := section.GetCurrentActivities(startTime, currentTime)
		if len(sectionActivities) > 0 {
			currentActivities[section.Name] = sectionActivities
		}
	}
	return currentActivities
}
