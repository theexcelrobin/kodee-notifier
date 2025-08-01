package main

import (
	kodeeLogger "github.com/theexcelrobin/kodee-notifier/internal/logger"
)

func main() {
	logger, err := kodeeLogger.NewLogger()
	if err != nil {
		panic(err)
	}

	defer logger.LogFile.Close()
}
