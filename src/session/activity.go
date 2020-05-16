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

func (act activity) Id() string {
	return act.id
}

func (act activity) Activity() string {
	return act.activity
}

func (act activity) Done() bool {
	return act.done
}

func (act *activity) SetDone() error {
	act.done = true
	return nil
}
