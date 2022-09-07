package utils

import (
	"log"
	"os"
)

// GetEnvValue ...
func GetEnvValue(key string) string {
	return os.Getenv(key)
}

// HACK move this function to infrastructure/logger/base.go in order to make the output more readable.
func RecoverPanic() {
	err := recover()
	if err != nil {
		log.Println("PANIC:", err)
	}
}
