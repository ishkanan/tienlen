package utils

import (
	"fmt"
	"time"
)

func now() string {
	return time.Now().Format("2-Jan 15:04:05")
}

// LogDebug prints a DEBUG message to stdout
func LogDebug(message string, params ...interface{}) {
	fmt.Printf("[%s] DEBUG  %s\n", now(), fmt.Sprintf(message, params...))
}

// LogInfo prints a INFO message to stdout
func LogInfo(message string, params ...interface{}) {
	fmt.Printf("[%s] INFO   %s\n", now(), fmt.Sprintf(message, params...))
}
