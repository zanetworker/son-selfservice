package configuration

import (
	"encoding/json"
	"os"

	log "github.com/Sirupsen/logrus"
)

//AppConfig is the configuration struct
var AppConfig Configuration

//LoadAppConfig read configuration from config file
func LoadAppConfig() {
	file, err := os.Open("configuration/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = Configuration{}
	err = decoder.Decode(&AppConfig)
	log.Infof("[ BootStrapper ]: DB= %s:%s",
		AppConfig.DBHostIP,
		AppConfig.DBHostPort,
	)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
