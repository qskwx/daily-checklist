package session

type category struct {
	name       string
	activities activities
	prefix     string
}

func categoryFabric(name string, prefix string, acts []string) category {
	return category{
		name:       name,
		activities: activitiesFabric(prefix, acts),
		prefix:     prefix,
	}
}

func (cat category) Name() string {
	return cat.name
}

func (cat category) Activities() activities {
	return cat.activities
}

func (cat *category) setDone(activityId string) error {
	return cat.activities.setDone(activityId)
}
