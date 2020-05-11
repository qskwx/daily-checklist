package session

import (
	userconfig "daily-checklist/src/user_config"
	"encoding/json"
	"fmt"
	"time"
)

type session struct {
	user       userconfig.UserConfig
	categories categories
}

// TODO: check username as input

func sessionFabric(username string) session {
	user, _ := userconfig.UserConfigFabric(username)
	ss := session{
		user:       user,
		categories: categoriesFabric(user.GetCurrentActivities(time.Now()))}
	return ss
}

func (ss *session) setDone(actID string) error {
	if err := ss.categories.setDone(actID); err != nil {
		return err
	}
	return nil
}

func (ss session) show() ([]byte, error) {
	byteArray, err := json.Marshal(ss.categories)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Unable to extract session info: %s", err))
	}
	return byteArray, nil
}

func IsExpired() int {
	return 0
}

func toCharStr(i int) string {
	return string('a' + i)
}
