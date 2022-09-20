package main

import (
	"github.com/sirupsen/logrus"
	"testing"
)

/*

Field
Sometimes you need to embed any information to the log such as the information of the user who login.
you can use logger.WithField(key, value) before logging to embed information to the log.

Fields
If you have many information to embed to log, you can use Fields instead of Field to embed more than one information.
You can use logger.WithFields(Logrus.Fields) . Fields is an alias for a type map[string]interface{} .

*/
func TestFeilds(t *testing.T) {

	logger := logrus.New()

	logger.WithField("username", "dj").Info("hello world")

	logger.WithFields(logrus.Fields{

		"username":  "dev",
		"full name": "dev ahsdz",
	}).Info("hi")

}
