package minictl

import (
	"com.github.mcruzdev.miniplatform.minictl/create"
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "minictl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run inside minictl")
	},
}

func Execute() {
	rootCmd.AddCommand(create.NewCreateCommand())

	rootCmd.Execute()
}
