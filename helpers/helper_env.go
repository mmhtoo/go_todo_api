package helpers

import (
	"errors"
	"fmt"
	"os"
)

func LoadFromEnv(key string, needThrow bool, fallback string) (string, error) {
	loadedValue := os.Getenv(key)
	fmt.Println(key, " = ", loadedValue)
	if loadedValue == "" {
		if needThrow {
			return "", errors.New(fmt.Sprintf("Failed to load env with key %v", key))
		}
		return fallback, nil
	}
	return loadedValue, nil
}
