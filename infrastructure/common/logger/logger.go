package logger

import (
	"fmt"

	"star-wars-api/configs"

	"go.uber.org/zap"
)

func NewLogger(cfg *configs.Configs) (*zap.Logger, error) {
	fileName := fmt.Sprintf("logs/%s.log", cfg.AppPrefix)
	logger := zap.NewProductionConfig()
	logger.OutputPaths = []string{
		fileName,
		"stderr",
	}
	return logger.Build()
}
