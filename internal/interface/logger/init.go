package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type LoggingVar struct {
	Log *zap.Logger
}

func NewLoggingVar() *LoggingVar {
	return &LoggingVar{
		Log: zap.NewNop(),
	}
}

func (l LoggingVar) Initialize(level string) error {
	_, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return fmt.Errorf("ошибка при прасинге уровня логера: %w", err)
	}
	return nil
}

func (l LoggingVar) Info(msg string, opt ...any) {
	l.Log.Info(fmt.Sprintf(msg, opt...))
}

func (l LoggingVar) Error(msg string, opt ...any) {
	l.Log.Error(fmt.Sprintf(msg, opt...))
}

func (l LoggingVar) Debug(msg string, opt ...any) {
	l.Log.Debug(fmt.Sprintf(msg, opt...))
}

func (l LoggingVar) Fatal(msg string, opt ...any) {
	l.Log.Fatal(fmt.Sprintf(msg, opt...))
}
