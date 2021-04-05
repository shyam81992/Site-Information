package config

import (
	"fmt"
	"os"
)

//AppConfig containes the config values for the app
var AppConfig = make(map[string]string)

//Rabbitmq containes the db credentials
var RabbitConfig = make(map[string]string)

func getEnvValue(key string, strict ...bool) (value string) {
	if len(strict) > 0 && strict[0] {
		if len(os.Getenv(key)) > 0 {
			value = os.Getenv(key)
		} else {
			fmt.Println("Environment not completely defined!! - execute environment config file!!")
			fmt.Println("Parameter - " + key + " is missing!")
			os.Exit(1)
		}
	} else {
		value = os.Getenv(key)
	}
	return value
}

//LoadConfig func loads the config from environment
func LoadConfig() {
	AppConfig["host"] = getEnvValue("HOST", true)
	AppConfig["port"] = getEnvValue("PORT", true)

	RabbitConfig["uri"] = getEnvValue("RMQ_URI", true)
	RabbitConfig["queuename"] = getEnvValue("RMQ_QUEUENAME", true)

}
