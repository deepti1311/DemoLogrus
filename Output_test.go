package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

/*

If you run the above code for the first time,
it would create a new file with name application.log .
If application.log already exist,
it would write log into application.log .

*/

func TestOutput(t *testing.T) {

	logger := logrus.New()

	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("info logging")
	logger.Warn("warning logging")
	logger.Error("Error logging")

}
