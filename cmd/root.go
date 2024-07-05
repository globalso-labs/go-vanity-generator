package cmd

import (
	"os"

	"github.com/gsols/go-logger"
	"github.com/nickaguilarh/credentials/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Short: "This service is responsible for receiving the messages from the devices and handling them.",
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		verbosity, _ := cmd.Flags().GetCount("verbosity")

		opt := logger.WithVerbosity(verbosity)
		if verbosity < int(logger.DefaultLevel) {
			opt = logger.WithVerbosity(int(logger.DefaultLevel))
		}

		logger.WithOptions(opt)

		ctx := logger.With().Int("pid", os.Getpid()).Logger().WithContext(cmd.Context())
		cmd.SetContext(ctx)

		config.Bootstrap(cmd.Context())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(InitCommand())

	rootCmd.PersistentFlags().CountP("verbosity", "v", "verbosity level")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
