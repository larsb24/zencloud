package environment

import (
	"errors"
	"log"
	"os"
)

func GetEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", errors.New("environment variable not set")
}

func GetPort() string {
	port, err := GetEnv("ZENCLOUD_BACKEND_PORT")
	if err != nil {
		log.Default().Println("Environment variable ZENCLOUD_BACKEND_PORT not set, using default port 8080")
		return "8080"
	}
	return port
}

func GetStorageLocation() string {
	storagePath, err := GetEnv("ZENCLOUD_BACKEND_STORAGE_PATH")
	if err != nil {
		log.Default().Println("Environment variable ZENCLOUD_BACKEND_STORAGE_PATH not set, using default path /tmp/")
		return "/tmp/"
	}
	return storagePath
}
