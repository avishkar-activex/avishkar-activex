package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func Init() {
	log.Infof("reading config file details")

	viper.SetConfigType("yaml")
	viper.Set("Verbose", true)
	// Note that config name does not include extension
	viper.SetConfigName("svc_config")

	baseDirPath, err := os.Getwd()
	if err != nil {
		log.Errorf("failed to get current directory path")
	}
	log.Infof("current directory path %v", baseDirPath)

	// Look for config file in source directory.
	// Requires GOPATH to be set even if using the
	// default location.
	configPath := fmt.Sprintf("%s/config", baseDirPath)
	viper.AddConfigPath(configPath)

	err = ParseConfig()
	if err != nil {
		log.Fatalf("fatal error, problem parsing config: %v", err)
	}
}

func ParseConfig() error {

	// Read config
	err := viper.ReadInConfig()
	if err != nil {
		log.Infof("error reading config file: %v", err)
		return err
	}

	return nil

}
