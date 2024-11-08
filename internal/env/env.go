package env

import (
	"log"
	"os"
	"strconv"
)

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}

func GetString(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	log.Printf(val)
	if !ok {
		return fallback
	}

	return val
}
