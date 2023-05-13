package create

import (
	"com.github.mcruzdev.miniplatform.minictl/internal/github/repo"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var applicationName string

func NewCreateApplicationCommand() *cobra.Command {
	var createApplicationCmd = &cobra.Command{
		Use: "application",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(fmt.Sprintf("Creating application with name '%s'", applicationName))
			newRepo := repo.NewRepo()
			newRepo.Create(applicationName)
		},
	}

	createApplicationCmd.PersistentFlags().StringVar(&applicationName, "name", "", "--name=application-name")
	_ = createApplicationCmd.MarkPersistentFlagRequired("name")

	return createApplicationCmd
}
