package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	LogFile *os.File
	Logger  *zerolog.Logger
}

func NewLogger() (*Logger, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logFile, err := os.OpenFile(os.Getenv("LOG_FILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := zerolog.New(logFile).With().Timestamp().Logger()

	return &Logger{
		LogFile: logFile,
		Logger:  &logger,
	}, nil
}
