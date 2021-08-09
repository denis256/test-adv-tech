package main

import (
	"os"
)

func getenv(name string, defaultValue string) string {
	v := os.Getenv(name)
	if v == "" {
		return defaultValue
	}

	return v
}
