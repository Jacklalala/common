package logger_test

import (
	"errors"
	"testing"

	"kchain.com/framework/logger"
)

func TestAll(t *testing.T) {
	logger.Debug("debug logging")
	logger.Info("debug logging", "dsfds")
	logger.Info(&struct{ Metadata string }{Metadata: "shit happen"})
	logger.Errorf("error %v", errors.New("shit happen"))
}
