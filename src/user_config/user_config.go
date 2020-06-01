package userconfig

import (
	"daily-checklist/src/user_config/sections"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// UserConfig is a type, which contains user-specific info
type UserConfig struct {
	userConfig
}

type userConfig struct {
	Name     string
	Start    string
	Sections sections.Sections
}

// Fabric create new UserConfig instance by given config.json
func Fabric(filename string) (UserConfig, error) {
	var user UserConfig

	configFile, err := os.Open(filename)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err)
		return user, err
	}
	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	err = json.Unmarshal(byteValue, &user)

	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

// GetCurrentActivities extract all activites related for particular day
func (conf UserConfig) GetCurrentActivities(currentTime time.Time) map[string][]string {
	startTime, _ := time.Parse("2006-02-01 15:04", conf.Start)
	currentActivities := conf.Sections.GetCurrentActivities(startTime, currentTime)
	return currentActivities
}
