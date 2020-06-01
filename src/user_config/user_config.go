package userconfig

import (
	"daily-checklist/src/user_config/sections"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type UserConfig struct {
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
	json.Unmarshal(byteValue, &user)
	return user, nil
}

// GetCurrentActivities extract all activites related for particular day
func (conf UserConfig) GetCurrentActivities(currentTime time.Time) map[string][]string {
	startTime, _ := time.Parse("2006-02-01 15:04", conf.Start)
	currentActivities := conf.Sections.GetCurrentActivities(startTime, currentTime)
	return currentActivities
}
