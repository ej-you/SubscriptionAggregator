// Package logger provides Init function to setup global logrus logger.
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// JSONFormatterUTC is the logrus.JSONFormatter wrapper with time in UTC.
type JSONFormatterUTC struct {
	logrus.JSONFormatter
}

// Format implements logrus.Formatter and sets time to UTC.
func (f *JSONFormatterUTC) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return f.JSONFormatter.Format(e)
}

// TextFormatterUTC is the logrus.TextFormatter wrapper with time in UTC.
type TextFormatterUTC struct {
	logrus.TextFormatter
}

// Format implements logrus.Formatter and sets time to UTC.
func (f *TextFormatterUTC) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return f.TextFormatter.Format(e)
}

// Init sets up main logger for application.
func Init() {
	logrus.SetOutput(os.Stderr)
	logrus.SetFormatter(&TextFormatterUTC{
		logrus.TextFormatter{FullTimestamp: true},
	})
}
