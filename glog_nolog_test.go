package glog3

import (
	"testing"
)

func TestNoLog(t *testing.T) {
	logging.toStderr = false
	logging.verbosity.Set("2")
	logging.nolog = true
	Info("test -- should not create log or files")
	logging.nolog = false
}
