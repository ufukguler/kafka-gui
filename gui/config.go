package gui

import (
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const (
	configType = "yml"
	configName = ".kafka-gui"
	configFile = configName + "." + configType
)

var cfgFile string

func InitConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Panicf("Error %v", err)
			return
		}
		viper.AddConfigPath(home)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		createConfigFile()
	}
}

func createConfigFile() {
	log.Info("Creating config file")
	config := []byte(`
WIDTH: 1366
HEIGHT: 800
`)
	home, err := homedir.Dir()
	if err != nil {
		log.Panicf("Error %v", err)
	}

	file, err := os.Create(home + "\\" + configFile)
	if err != nil {
		log.Panicf("Error %v", err)
	}

	_, err = file.Write(config)
	if err != nil {
		log.Panicf("Error %v", err)
	}

	err = file.Close()
	if err != nil {
		log.Panicf("Error %v", err)
	}
}

func viperSetConfig(key, value string) {
	viper.Set(key, value)
	err := viper.WriteConfig()

	if err != nil {
		log.Panicf("Error %v", err)
	}
}
