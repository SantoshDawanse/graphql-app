package logger

import (
	log "github.com/sirupsen/logrus"
)

func InitializeLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	// Add the calling method as a field
	log.SetReportCaller(true)
	// Will log anything that is info or above (warn, error, fatal, panic). Default.
	log.SetLevel(log.InfoLevel)
}
