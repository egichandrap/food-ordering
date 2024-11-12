package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Function for get timestamp on format YYYYMMDDHHMMSS
func GenerateTimestamp() string {
	return time.Now().Format("20060102150405")
}

// Function for create string random with length uniq
func GenerateRandomString(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
