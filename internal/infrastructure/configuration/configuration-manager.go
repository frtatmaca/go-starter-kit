package configuration

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

type IConfigurationManager interface {
	GetServerConfig() ServerConfig
	GetMongoDbConfig() MongoDbConfig
	GetLogConfig() LogConfig
}

const (
	serverPort = ":8080"
	configType = "yaml"
)

var instance *ConfigurationManager

type ConfigurationManager struct {
	applicationConfig ApplicationConfig
}

func NewConfigurationManager(configPath string, configName string) IConfigurationManager {
	if instance != nil {
		return instance
	}

	env := os.Getenv("ACTIVE_PROFILE")
	if env == "" {
		log.Print("**** ACTIVE_PROFILE is empty, default it will be used as 'dev' ****")
		env = "dev"
	}

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetTypeByDefaultValue(true)
	viper.SetDefault("server.port", serverPort)
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		readConf(env, configPath, configName)
	})

	instance := readConf(env, configPath, configName)

	return &ConfigurationManager{
		applicationConfig: instance,
	}
}

func (configurationManager *ConfigurationManager) GetServerConfig() ServerConfig {
	return configurationManager.applicationConfig.Server
}

func (configurationManager ConfigurationManager) GetMongoDbConfig() MongoDbConfig {
	return configurationManager.applicationConfig.MongoDb
}

func (configurationManager ConfigurationManager) GetLogConfig() LogConfig {
	return configurationManager.applicationConfig.Log
}

func readConf(env string, configPath string, configName string) ApplicationConfig {
	readConfigErr := viper.ReadInConfig()
	if readConfigErr != nil {
		log.Panicf("Couldn't load application configuration, cannot start. Error details: %s", readConfigErr.Error())
	}
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetTypeByDefaultValue(true)
	mergeConfigErr := viper.MergeInConfig()
	if mergeConfigErr != nil {
		log.Panicf("Couldn't load application configuration, cannot start. Error details: %s", mergeConfigErr.Error())
	}
	var conf ApplicationConfig
	c := viper.Sub(env)
	unMarshalErr := c.Unmarshal(&conf)
	unMarshalSubErr := c.Unmarshal(&conf)

	if unMarshalErr != nil {
		log.Panicf("Configuration cannot deserialize. Terminating. Error details: %s", unMarshalErr.Error())
	}
	if unMarshalSubErr != nil {
		log.Panicf("Configuration cannot deserialize. Terminating. Error details: %s", unMarshalSubErr.Error())
	}

	return conf
}
