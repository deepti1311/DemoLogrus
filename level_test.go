package main

import (
	"github.com/sirupsen/logrus"
	"testing"
)

//running a unit test for logging level

/*

Level
Level is a determination of priority or type of an event.
Level start from the lowest to the highest level.
Logrus is supporting many levels.
Please see the table below for the list of Logging Level.


*/

func TestLevel(t *testing.T) {

	logger := logrus.New()

	logger.Trace("this is a Trace")
	logger.Debug("this is a Debug")
	logger.Info("this is a Info")
	logger.Warn("this is warn")
	logger.Error("this is Error")

}

/* you may found that trace & debug are not printed in the console.
It happens because, by default, there is only Info priority and higher will be logged.
But you can set the logging level by using logger.SetLevel() .

*/
