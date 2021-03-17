package main

import (
	"os"
	"seed_phrase_generator/internal/logger"
	"seed_phrase_generator/internal/seedgenerator"
	"seed_phrase_generator/internal/utils/ltrswitcher"
	"seed_phrase_generator/internal/utils/txtparser"

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

	gen := seedgenerator.NewSeedGen(path, txtparser.InitParser(), ltrswitcher.NewSwitcher())
	gen.SelectSrc()
	gen.SelectNumber()
	gen.ShowChapters()
}
