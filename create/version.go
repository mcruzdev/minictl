package create

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var version string

func NewCreateVersionCommand() *cobra.Command {
	createVersionCommand := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(fmt.Sprintf("Creating a new version with semver number '%s'", version))
		},
	}

	createVersionCommand.PersistentFlags().StringVar(&version, "number", "", "--version=0.1.0")
	_ = createVersionCommand.MarkPersistentFlagRequired("number")

	return createVersionCommand
}
