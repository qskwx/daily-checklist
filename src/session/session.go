package session

import (
	userconfig "daily-checklist/src/user_config"
	"fmt"
	"time"
)

type Session struct {
	user         userconfig.UserConfig
	categories   categories
	creationTime time.Time
}

func SessionFabric(username string) (Session, error) {
	username = fmt.Sprintf("configs/%s.json", username) // TODO: reformat
	user, err := userconfig.Fabric(username)
	if err != nil {
		return Session{}, err
	}
	creationTime := time.Now()
	return Session{
		user:         user,
		categories:   categoriesFabric(user.GetCurrentActivities(creationTime)),
		creationTime: creationTime}, nil
}

func (ss *Session) SetDone(actID string) error {
	return ss.categories.setDone(actID)
}

func (ss Session) Categories() categories {
	return ss.categories
}

func (ss Session) IsActual(now time.Time) bool {
	passed := now.Sub(ss.creationTime).Hours()
	return passed < 24
}
