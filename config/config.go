package config

import (
	"os"
	"strconv"
)

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

func GetTbkAuthInfo() (appKey, secret string) {
	appKey = os.Getenv("AppKey")
	secret = os.Getenv("Secret")
	return
}

func GetAdzoneID() int64 {
	v, _ := strconv.ParseInt(os.Getenv("AdzoneID"), 10, 64)
	return v
}

func GetPddAuthInfo() (clientID, secret string) {
	clientID = os.Getenv("PddClientID")
	secret = os.Getenv("PddSecret")
	return
}

func GetPddPID() string {
	return "13449638_213518026"
}

func GetAuthToken() string {
	return os.Getenv("AuthToken")
}
