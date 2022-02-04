package main

import (
	log "github.com/sirupsen/logrus"
	. "kafka-gui/gui"
	"time"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
	InitConfig()
	Run()
}
