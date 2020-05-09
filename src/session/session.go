package session

import (
	userconfig "daily-checklist/src/user_config"
)

type session struct {
	user       userconfig.UserConfig
	categories []category
}

type category struct {
	name        string
	prefix      string
	acitivities []activity
}

type activity struct {
	id       string
	activity string
	done     bool
}

func IsExpired() int {
	return 0
}
