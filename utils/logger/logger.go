package utils

import (
	"log"
	"time"
)

const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

func Info(message string) {
	Log(INFO, message)
}

func Warn(message string) {
	Log(WARN, message)
}

func Err(message string) {
	Log(ERROR, message)
}

func Log(level string, message string) {
	timestamp := time.Now().String()
	log.Printf("%s [%s]: %s", level, timestamp, message)
}
