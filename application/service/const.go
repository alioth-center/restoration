package service

import (
	"github.com/alioth-center/infrastructure/logger"
)

var (
	restorationIdentificationKey = "restoration_identification"

	levels = map[logger.Level]struct{}{
		logger.LevelDebug: {},
		logger.LevelInfo:  {},
		logger.LevelWarn:  {},
		logger.LevelError: {},
		logger.LevelFatal: {},
		logger.LevelPanic: {},
	}
)
