package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.globalso.dev/x/tools/vanity/internal/constants/cmd"
	"go.globalso.dev/x/tools/vanity/internal/infraestructure/generate"
)

// generateCmd represents the execute command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the static files for the vanity URL server.",
	Long: `
Generate the static files for the vanity URL server. This command
will generate the necessary files for the vanity URL server to
function properly. It will not overwrite existing files unless
the force flag is set.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return generate.Execute(cmd.Context())
	},
}

func GenerateCommand() *cobra.Command {
	generateCmd.Flags().StringP("output", "o", "dist", "output directory for the generated files")
	_ = viper.BindPFlag(cmd.GeneratorOutput, generateCmd.Flags().Lookup("output"))

	generateCmd.Flags().BoolP("clean", "c", false, "clean the existing files before generating")
	_ = viper.BindPFlag(cmd.GeneratorClean, generateCmd.Flags().Lookup("clean"))

	return generateCmd
}
