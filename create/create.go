package create

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewCreateCommand handles minictl create command.
func NewCreateCommand() *cobra.Command {
	var createCmd = &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running create command")
		},
	}

	createCmd.AddCommand(NewCreateApplicationCommand())
	return createCmd
}
