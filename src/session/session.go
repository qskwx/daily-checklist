package session

import (
	userconfig "daily-checklist/src/user_config"
	"fmt"
	"time"
)

type Session struct {
	user       userconfig.UserConfig
	categories categories
}

// TODO: check username as input

func SessionFabric(username string) Session {
	username = fmt.Sprintf("src/user_config/configs/%s.json", username) // TODO: reformat this!
	user, _ := userconfig.UserConfigFabric(username)
	ss := Session{
		user:       user,
		categories: categoriesFabric(user.GetCurrentActivities(time.Now()))}
	return ss
}

func (ss *Session) SetDone(actID string) error {
	return ss.categories.setDone(actID)
}

func (ss Session) Categories() categories {
	return ss.categories
}
