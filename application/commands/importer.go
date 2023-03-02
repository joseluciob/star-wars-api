package commands

import (
	"log"
	"star-wars-api/configs"
	"star-wars-api/domain/service"
	"star-wars-api/infrastructure/persistence"
	"star-wars-api/infrastructure/provider"

	"github.com/spf13/cobra"
)

func ImportPlanetsCommand(cmd *cobra.Command, _ []string) error {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal(err)
	}

	repos, err := persistence.NewRepositories(&cfg.DB)
	if err != nil {
		panic(err)
	}
	defer repos.Close()
	repos.Automigrate()

	provider, err := provider.New(cfg)
	if err != nil {
		return err
	}

	film := service.NewFilmService(repos, provider)
	film.Import(cmd.Context())

	planet := service.NewPlanetService(repos, provider)
	planet.Import(cmd.Context())

	return nil
}
