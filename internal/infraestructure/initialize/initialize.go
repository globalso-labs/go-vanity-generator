/*
 * credentials
 * initialize.go
 * This file is part of credentials.
 * Copyright (c) 2024.
 * Last modified at Tue, 2 Jul 2024 21:49:26 -0500 by nick.
 *
 * DISCLAIMER: This software is provided "as is" without warranty of any kind, either expressed or implied. The entire
 * risk as to the quality and performance of the software is with you. In no event will the author be liable for any
 * damages, including any constants, special, incidental, or consequential damages arising out of the use or inability
 * to use the software (that includes, but not limited to, loss of data, data being rendered inaccurate, or losses
 * sustained by you or third parties, or a failure of the software to operate with any other programs), even if the
 * author has been advised of the possibility of such damages.
 * If a license file is provided with this software, all use of this software is governed by the terms and conditions
 * set forth in that license file. If no license file is provided, no rights are granted to use, modify, distribute,
 * or otherwise exploit this software.
 */

package initialize

import (
	"context"
	"io"
	"os"
	"runtime"

	"github.com/gsols/go-logger"
	"github.com/spf13/viper"
	"go.globalso.dev/x/tools/vanity/internal"
	"go.globalso.dev/x/tools/vanity/internal/data"
)

// Execute runs the application and handles the shutdown signal.
func Execute(ctx context.Context) error {
	ctx = logger.Ctx(ctx).With().Str("module", "initialize").Logger().WithContext(ctx)

	logger.Ctx(ctx).Info().Msgf("Running version %s on %s/%s with %d CPUs and %d GOMAXPROCS.",
		internal.Version, runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), runtime.GOMAXPROCS(0),
	)
	// Check if the configuration file exists or if the force flag is set to overwrite it.
	_, err := os.Stat(internal.ConfigPath)
	forceOverwrite := viper.GetBool("force")
	if err == nil && !forceOverwrite {
		logger.Ctx(ctx).Warn().Msgf("Configuration file already exists at %s. Skipping...", internal.ConfigPath)
		return nil
	}

	// If the force flag is set, overwrite the configuration file.
	if forceOverwrite {
		logger.Ctx(ctx).Warn().Msgf("Force flag set. Overwriting configuration file at %s.", internal.ConfigPath)
	}

	// Write the configuration file to disk.
	err = writeConfigFile(ctx, internal.ConfigPath)
	if err != nil {
		logger.Ctx(ctx).Error().Err(err).Msg("Failed to write configuration file")
		return err
	}

	return nil
}

func writeConfigFile(ctx context.Context, configPath string) error {
	logger.Ctx(ctx).Info().Msgf("Writing configuration file to %s.", configPath)

	// Write the configuration file to disk.
	config, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer config.Close()

	input, err := data.ExampleFile.Open(data.ExampleFilePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(config, input)
	if err != nil {
		return err
	}

	return nil
}
