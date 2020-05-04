package userconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type periodicity struct {
	Metrics     string
	Denominator int
	Addendum    int
}

type borders struct {
	Type        string
	LeftBorder  int
	RightBorder int
}

type growFunction struct {
	Type        string
	Coefficient int
}

type grow struct {
	Borders      borders
	GrowFunction growFunction
}

type activity struct {
	Name        string
	Periodicity periodicity
	Grow        grow
	Group       int
}

type section struct {
	SectionName string `json:"section"`
	Activities  []activity
}

type tasks []section

// UserConfig is a type, which contains user-specific info
type UserConfig struct {
	Name  string
	Tasks tasks
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
	fmt.Println(user.Tasks[0])

	return user, nil
}
