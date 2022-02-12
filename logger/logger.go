package logger

import (
	log "github.com/sirupsen/logrus"
)

func Info(args interface{}) {
	log.Info(args)
}
