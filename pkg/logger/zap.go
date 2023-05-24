package logger

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Initialize() (logger *zap.Logger, err error) {
	if viper.GetString("server.mode") != "local" {
		logger, err = zap.NewProduction(
			zap.Hooks(func(entry zapcore.Entry) error {
				if entry.Level == zapcore.ErrorLevel {
					defer sentry.Flush(2 * time.Second)
					sentry.CaptureMessage(fmt.Sprintf("%s, Line No: %d :: %s", entry.Caller.File, entry.Caller.Line, entry.Message))
				}
				return nil
			}))
		if err != nil {
			return nil, err
		}
		return logger, nil

	} else {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
		return logger, nil
	}
}
