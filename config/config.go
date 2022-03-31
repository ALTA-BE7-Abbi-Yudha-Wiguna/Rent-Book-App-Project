package config

import (
	"os"
	"sync"
	"fmt"
)

type AppConfig struct {
	Port     string `yaml:"port"`
	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {

	var defaultConfig AppConfig
	defaultConfig.Port = os.Getenv("PORT")
	defaultConfig.Database.Driver = os.Getenv("DB_DRIVER")
	defaultConfig.Database.Name = os.Getenv("DB_NAME")
	defaultConfig.Database.Address = os.Getenv("ADDRESS")
	defaultConfig.Database.Port = os.Getenv("DB_PORT")
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")
	fmt.Println(defaultConfig)
	return &defaultConfig
}
