package main

import (
	"context"
	"log"
	"star-wars-api/application/commands"
	"star-wars-api/configs"
	logg "star-wars-api/infrastructure/common/logger"

	"github.com/spf13/cobra"
)

func NewImporterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Interface to import data from Swapi.dev API",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "planets",
		Short: "Import all planets",
		RunE:  commands.ImportPlanetsCommand,
	})

	return cmd
}

func main() {
	cfg := cfg()
	logger, err := logg.NewLogger(cfg)
	if err != nil {
		log.Panic("panic:", err)
	}
	defer logger.Sync()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmds := NewImporterCommand()
	if err := cmds.ExecuteContext(ctx); err != nil {
		logger.Fatal("error: ", logg.ErrorField(err))
	}
}

func cfg() *configs.Configs {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal("fatal error: ", err)
	}

	return cfg
}
