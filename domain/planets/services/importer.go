package services

import (
	"log"
	"star-wars-api/configs"
	"star-wars-api/infrastructure/planets/provider"

	"github.com/spf13/cobra"
)

func ImportPlanetsCommand(cmd *cobra.Command, _ []string) error {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal(err)
	}

	p, err := provider.New(cfg)
	if err != nil {
		return err
	}

	page := 1
	for {

		results := p.GetPlanets(cmd.Context(), page)
		if results.Next == "" {
			break
		}

		page = page + 1
	}

	return nil
}
