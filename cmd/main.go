package main

import (
	"os"
	"seed_phrase_generator/internal/logger"

	"go.uber.org/zap"
)

var (
	txtDir    = "text_files"
	appName   = "seed_generator"
	buildTime = "_dev"
	buildHash = "_dev"
)

func main() {
	logger.InitLogger()
	logger.Info("start app", zap.String("app", appName), zap.String("build", buildHash), zap.String("time", buildTime))
	path, err := os.Getwd()
	if err != nil {
		logger.Fatal("Can't locate current dir", err)
	}
	path += string(os.PathSeparator) + txtDir
	logger.Info("text storage path", zap.String("path", path))
}
