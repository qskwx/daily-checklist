package session

import (
	userconfig "daily-checklist/src/user_config"
	"fmt"
	"time"
)

type Session struct {
	user          userconfig.UserConfig
	categories    categories
	pastFromStart int
}

func SessionFabric(username string) (Session, error) {
	username = fmt.Sprintf("configs/%s.json", username) // TODO: reformat
	user, err := userconfig.Fabric(username)
	if err != nil {
		return Session{}, err
	}
	pastFromStart := daysPast(user.StartTime(), time.Now())
	return Session{
		user:          user,
		categories:    categoriesFabric(user.GetCurrentActivities(time.Now())),
		pastFromStart: pastFromStart}, nil
}

func (ss *Session) SetDone(actID string) error {
	return ss.categories.setDone(actID)
}

func (ss Session) Categories() categories {
	return ss.categories
}

func (ss Session) IsActual(now time.Time) bool {
	passed := daysPast(ss.user.StartTime(), time.Now())
	return passed > ss.pastFromStart
}

func daysPast(startTime time.Time, now time.Time) int {
	pastFromStart := int(now.Sub(startTime).Hours() / 24)
	return pastFromStart
}
