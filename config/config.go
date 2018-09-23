package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Configuration : includes the credential to talk with Emarsys API
type Configuration struct {
	User    string `json:"user"`
	Secret  string `json:"secret"`
	BaseURL string `json:"baseUrl"`
	EventID string `json:"eventId"`
	KeyID   int8   `json:"keyId"`
	MySQL   string `json:"msyql"`
}

// Items test shared config
var Items *Configuration

func init() {

	if Items != nil {
		return
	}

	basePath, err := os.Getwd()
	panicError(err)

	bts, err := ioutil.ReadFile(filepath.Join(basePath, "config", "config.json"))
	panicError(err)

	Items = new(Configuration)
	err = json.Unmarshal(bts, &Items)
	panicError(err)
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

// GetConfig : Get emarsys configuration
// includes the credential to talk with Emarsys API
func GetConfig() Configuration {
	return *Items
}
