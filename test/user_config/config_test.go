package config_test

import (
	userconfig "daily-checklist/internal/user_config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

func setupTestCase(t *testing.T, source string, config string) func(t *testing.T) {
	t.Log("setup test case")
	data, err := ioutil.ReadFile(source)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create(config)
	defer file.Close()
	if err != nil {
		t.Error(err)
	}
	file.Write(data)
	return func(t *testing.T) {
		t.Log("teardown test case")
		if err := os.Remove(config); err != nil {
			t.Error(err)
		}
	}
}

type suit struct {
	Time   string `json:"current_time"`
	Result map[string][]string
}

func setupTestSuits(t *testing.T, filename string, suits *[]suit) func(t *testing.T) {
	t.Log("setup test suits")
	cf, err := os.Open(filename)
	defer cf.Close()
	if err != nil {
		t.Error(err)
	}
	byteValue, err := ioutil.ReadAll(cf)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(byteValue, &suits)
	if err != nil {
		t.Error(err)
	}
	return func(t *testing.T) {
		t.Log("teardown test suits")
	}
}

func TestCurrentActivities(t *testing.T) {
	sourceName := "suits/current_acts/input.json"
	configName := "testconfig.json"
	teardownTestCase := setupTestCase(t, sourceName, configName)
	defer teardownTestCase(t)

	var suits []suit
	suitsFile := "suits/current_acts/output.json"
	teardownTestSuits := setupTestSuits(t, suitsFile, &suits)
	defer teardownTestSuits(t)

	config, err := userconfig.Fabric(configName)
	if err != nil {
		t.Errorf("Error construction fabric catched: %s", err)
		return
	}

	for idx := range suits {
		t.Run(fmt.Sprintf("suit#%d", idx), func(t *testing.T) {
			currentTime, err := time.Parse("2006-02-01 15:04", suits[idx].Time)
			if err != nil {
				t.Errorf("Error construction time: %s", err)
				return
			}
			currentActs := config.GetCurrentActivities(currentTime)
			if reflect.DeepEqual(suits[idx].Result, currentActs) != true {
				t.Error("Error in current activities")
				t.Logf("expect:\t%v", suits[idx].Result)
				t.Logf("   got:\t%v", currentActs)
				return
			}
			return
		})
	}
}

func TestFabricWithEmptyFile(t *testing.T) {
	configName := "testconfig.json"
	_, err := userconfig.Fabric(configName)
	if err == nil {
		t.Errorf("Error wasnt catched on unexisted config;")
	}
}
