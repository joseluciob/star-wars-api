package main

import (
	"log"
	"star-wars-api/application/handlers"
	"star-wars-api/application/handlers/middleware"
	"star-wars-api/configs"
	"star-wars-api/domain/service"
	logg "star-wars-api/infrastructure/common/logger"
	"star-wars-api/infrastructure/persistence"
	"star-wars-api/infrastructure/provider"

	_ "star-wars-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title           SWApi
// @version         1.0
// @description     This api integrates with the service https://swapi.dev

// @contact.name   José Barbosa
// @contact.email  joseluciobj@gmail.com

// @host      localhost:8190
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cfg := cfg()
	logger, err := logg.NewLogger(cfg)
	if err != nil {
		log.Panic("panic:", err)
	}
	defer logger.Sync()

	repos, err := persistence.NewRepositories(cfg)
	if err != nil {
		logger.Panic("panic:", logg.ErrorField(err))
	}
	defer repos.Close()
	repos.Automigrate()

	provider, _ := provider.New(cfg)
	planet := service.NewPlanetService(repos, provider)
	handler := handlers.NewPlanet(planet)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	apiV1 := r.Group("/api/v1")
	apiV1.GET("/planets", handler.GetAll)
	apiV1.GET("/planets/:id", middleware.IdentifierMiddleware(), handler.Get)
	apiV1.DELETE("/planets/:id", middleware.IdentifierMiddleware(), handler.Delete)
	apiV1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Info("App: ", logg.ErrorField(r.Run(":"+cfg.Port)))

}

func cfg() *configs.Configs {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal("fatal error: ", err)
	}

	return cfg
}
