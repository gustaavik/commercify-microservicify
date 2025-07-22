package common

import (
	"fmt"
	"os"
)

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	fmt.Println("Checking environment variable:", key, "Value:", value)

	if value == "" {
		return defaultValue
	}
	return value
}
