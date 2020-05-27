// Config allows to define application behaviour.
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	appConfig *AppConfig
)

// AppConfig holds config values to define application behaviour.
type AppConfig struct {
	SendInitialHooks      bool
	PollingInterval       uint
	MonitoredShazamCharts []string
	Tracklists            []Tracklist
}

// A WebhookTarget defines a target identified by an URL to receive webhooks.
type WebhookTarget struct {
	URL string
}

// A Tracklist defines a list of chartlists to receive tracks from and also
// a list of targets to receive webhooks when new tracks have been added.
type Tracklist struct {
	WatchedShazamCharts []string
	WebhookTargets      []WebhookTarget
}

// Get loads and returns the singleton app config.
func Get() *AppConfig {
	if appConfig == nil {
		json, err := loadConfigFile()

		if err != nil {
			log.Fatal(err)
		}

		config, err := load(json)

		if err != nil {
			log.Fatal(err)
		}

		appConfig = config
	}

	return appConfig
}

func load(configJson []byte) (*AppConfig, error) {
	config := &AppConfig{}

	err := json.Unmarshal(configJson, config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func loadConfigFile() ([]byte, error) {
	ex, err := os.Executable()

	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(filepath.Dir(ex), "app-config.json")
	return ioutil.ReadFile(configPath)
}
