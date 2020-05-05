package userconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// UserConfig is a type, which contains user-specific info
type UserConfig struct {
	Name  string
	Tasks tasks
	Start string
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
	fmt.Println(user.Tasks[0])

	return user, nil
}

type tasks []section

type section struct {
	SectionName string `json:"section"`
	Activities  []activity
}
