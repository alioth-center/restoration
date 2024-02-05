package engine

import "github.com/alioth-center/infrastructure/logger"

func NewLogger(outputFile string) logger.Logger {
	if outputFile == "" {
		outputFile = "./restoration_collection.log"
	}

	log := logger.NewLoggerWithConfig(logger.Config{
		Level:          string(logger.LevelDebug),
		Formatter:      "json",
		StdoutFilePath: outputFile,
		StderrFilePath: "/dev/null",
	})

	return log
}
