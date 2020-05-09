package session

type activity struct {
	id       string
	activity string
	done     bool
}

func activityFabric(id string, action string, done bool) activity {
	return activity{
		id:       id,
		activity: action,
		done:     done,
	}
}
