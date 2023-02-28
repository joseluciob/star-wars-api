package main

import (
	"context"
	"log"
	"star-wars-api/application/commands"

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmds := NewImporterCommand()
	if err := cmds.ExecuteContext(ctx); err != nil {
		log.Fatal(err)
	}
}
