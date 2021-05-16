package config

import "os"

func GetDefaultEnv(key, defaultStr string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		value = defaultStr
	}
	return value
}

func IsProdEnv() bool {
	return GetDefaultEnv("IN_DOCKER", "0") != "0"
}
