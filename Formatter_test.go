package main

import (
	"github.com/sirupsen/logrus"
	"testing"
)

/*

When a logger sends a data to output,
the log which sends will be formatted using Object Formatter.
Logrus has two formatters, there are TextFormatter & JsonFormatter.
By default the formatter in Logrus will be set as TextFormatter.
You can change formatter by using logger.SetFormatter() .

*/

func TestFormatter(t *testing.T) {

	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("info log")
	logger.Warn("warn log")
	logger.Error("error log")
}
