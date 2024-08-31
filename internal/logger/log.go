package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

var Log *log.Logger

func Init() {
	Log = log.NewWithOptions(
		os.Stderr,
		log.Options{
			ReportCaller:    true,
			ReportTimestamp: true,
		})

	// Set the default level to Info
	Log.SetLevel(log.InfoLevel)
}
