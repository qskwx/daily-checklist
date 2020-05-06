package userconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// UserConfig is a type, which contains user-specific info
type UserConfig struct {
	Name     string
	Sections sections
	Start    string
}

// NewUserConfig create new UserConfig instance by given config.json
func NewUserConfig(filename string) (UserConfig, error) {
	var user UserConfig

	configFile, err := os.Open(filename)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err)
		return user, err
	}

	fmt.Println("Successfully Opened ", configFile)

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(byteValue, &user)

	fmt.Printf("Name: '%s'\n", user.Name)
	fmt.Printf("Date: '%s'\n", user.Start)
	fmt.Println(user.Sections[0])

	return user, nil
}

func (conf UserConfig) GetCurrentActivities(currentTime time.Time) map[string][]string {
	startTime, _ := time.Parse("2006-02-01 15:04", conf.Start)
	currentActivities := make(map[string][]string)
	for _, section := range conf.Sections {
		sectionActivities := make([]string, 4)
		for _, act := range section.Activities {
			if act.IsActual(startTime, currentTime) {
				sectionActivities = append(sectionActivities, act.GetSummary(startTime, currentTime))
			}
		}
		if len(sectionActivities) > 0 {
			currentActivities[section.Name] = sectionActivities
		}
	}
	return currentActivities
}

type sections []section

type section struct {
	Name       string `json:"section"`
	Activities []activity
}
