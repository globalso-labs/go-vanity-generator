package cmd

import (
	"github.com/nickaguilarh/credentials/internal/infraestructure/initialize"

	"github.com/spf13/cobra"
)

// initCmd represents the execute command.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the directory with the configuration file",
	Long: `
Initialize the directory with the configuration file
and the necessary files to start the project. Does
not overwrite existing files.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return initialize.Execute(cmd.Context())
	},
}

func InitCommand() *cobra.Command {
	return initCmd
}
