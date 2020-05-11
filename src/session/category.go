package session

type category struct {
	name       string
	prefix     string
	activities activities
}

func categoryFabric(name string, prefix string, acts []string) category {
	return category{
		name:       name,
		prefix:     prefix,
		activities: activitiesFabric(prefix, acts),
	}
}
