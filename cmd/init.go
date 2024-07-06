package cmd

import (
	"github.com/spf13/viper"
	"go.globalso.dev/x/tools/vanity/internal/constants/cmd"
	"go.globalso.dev/x/tools/vanity/internal/infraestructure/initialize"

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
	initCmd.Flags().BoolP("force", "f", false, "force the initialization")
	_ = viper.BindPFlag(cmd.Force, initCmd.Flags().Lookup("force"))

	return initCmd
}
