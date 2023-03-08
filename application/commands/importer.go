package commands

import (
	"star-wars-api/configs"
	"star-wars-api/domain/service"
	"star-wars-api/infrastructure/persistence"
	"star-wars-api/infrastructure/provider"

	"github.com/spf13/cobra"
)

func ImportPlanetsCommand(cmd *cobra.Command, _ []string) error {
	cfg, err := configs.New()
	if err != nil {
		return err
	}

	repos, err := persistence.NewRepositories(cfg)
	if err != nil {
		return err
	}
	defer repos.Close()
	repos.Automigrate()

	provider, err := provider.New(cfg)
	if err != nil {
		return err
	}

	film := service.NewFilmService(repos, provider)
	err = film.Import(cmd.Context())
	if err != nil {
		return err
	}

	planet := service.NewPlanetService(repos, provider)
	err = planet.Import(cmd.Context())
	if err != nil {
		return err
	}

	return nil
}
