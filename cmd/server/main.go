package main

import (
	"log"

	"star-wars-api/application/handlers"
	"star-wars-api/configs"
	"star-wars-api/domain/service"
	"star-wars-api/infrastructure/common/logger"
	"star-wars-api/infrastructure/persistence"
	"star-wars-api/infrastructure/provider"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := cfg()
	logger, err := logger.NewLogger(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer logger.Sync()

	repos, err := persistence.NewRepositories(&cfg.DB)
	if err != nil {
		panic(err)
	}
	defer repos.Close()
	repos.Automigrate()

	provider, _ := provider.New(cfg)
	planet := service.NewPlanetService(repos, provider)
	handler := handlers.NewPlanet(planet)

	r := gin.Default()
	r.GET("/planets", handler.GetAll)
	r.GET("/planets/:id", handler.Get)
	r.DELETE("/planets/:id", handler.Delete)

	log.Fatal(r.Run(":" + cfg.Port))

}

func cfg() *configs.Configs {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
