package logger

import (
	"errors"
	"testing"
)

func TestAll(t *testing.T) {
	InitLevel("test",5)
	Debug("debug logging")
	Info("debug logging", "dsfds")
	Info(&struct{ Metadata string }{Metadata: "metadata"})
	Errorf("error %v", errors.New("error test"))
}
