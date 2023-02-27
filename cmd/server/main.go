package api

import (
	"log"

	"star-wars-api/configs"
	"star-wars-api/infrastructure/common/logger"
)

func main() {
	cfg := cfg()
	logger, err := logger.NewLogger(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer logger.Sync()

}

func cfg() *configs.Configs {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
