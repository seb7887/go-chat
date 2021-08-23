package config

import (
	"os"
	"reflect"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	ServerPort int `env:"SERVER_PORT" default:"8080" json:"serverPort"`
}

var config *Configuration
var once sync.Once

// Return configuration
func GetConfig() *Configuration {
	once.Do(func() {
		setup()
	})
	return config
}

func setup() {
	// Try to load env vars from .env file
	loadEnv()
	config = &Configuration{}

	// Parse env vars and set field of config struct
	configReflect := reflect.ValueOf(config).Elem()
	err := loadConfig(configReflect, configReflect.Type())

	if err != nil {
		log.Error("error in reading configuration")
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Debugf("Could not load .env file, using environment variables. Error: %v", err)
	}
}

func loadConfigField(field reflect.Value, fieldDef reflect.StructField) error {
	var err error
	configField := fieldDef.Tag.Get("env")
	defaultValue := fieldDef.Tag.Get("default")
	configValue := os.Getenv(configField)

	if len(configValue) == 0 {
		configValue = defaultValue
	}

	switch field.Type().Kind() {
	case reflect.String:
		field.SetString(configValue)
		log.Debug("Loaded configuration")
	case reflect.Int:
		intValue, err := strconv.Atoi(configValue)
		if err != nil {
			log.Error("Invalid configuration")
		} else {
			field.SetInt(int64(intValue))
			log.Debug("Loaded configuration")
		}
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(configValue)
		if err != nil {
			log.Error("Invalid configuration")
		} else {
			field.SetBool(boolValue)
			log.Debug("Loaded configuration")
		}
	}

	return err
}

func loadConfig(configValue reflect.Value, configValueType reflect.Type) error {
	var err error
	for i := 0; i < configValue.NumField(); i++ {
		field := configValue.Field(i)
		err = loadConfigField(field, configValueType.Field(i))
	}
	return err
}
