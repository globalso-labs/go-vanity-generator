/*
 * credentials
 * reader.go
 * This file is part of credentials.
 * Copyright (c) 2024.
 * Last modified at Fri, 5 Jul 2024 02:55:45 -0500 by nick.
 *
 * DISCLAIMER: This software is provided "as is" without warranty of any kind, either expressed or implied. The entire
 * risk as to the quality and performance of the software is with you. In no event will the author be liable for any
 * damages, including any general, special, incidental, or consequential damages arising out of the use or inability
 * to use the software (that includes, but not limited to, loss of data, data being rendered inaccurate, or losses
 * sustained by you or third parties, or a failure of the software to operate with any other programs), even if the
 * author has been advised of the possibility of such damages.
 * If a license file is provided with this software, all use of this software is governed by the terms and conditions
 * set forth in that license file. If no license file is provided, no rights are granted to use, modify, distribute,
 * or otherwise exploit this software.
 */

package config

import (
	"context"
	"strings"

	"github.com/gsols/go-logger"
	"github.com/nickaguilarh/credentials/pkg/errors"
	"github.com/spf13/viper"
)

// ReadToStruct is a function that attempts to unmarshal the configuration data
// from Viper into the config struct. If it fails, it logs a fatal error.
// If it succeeds, it logs a trace message indicating success.
func ReadToStruct(ctx context.Context) {
	if err := viper.Unmarshal(config); err != nil {
		logger.Fatal().Err(err).Msg("Unable to decode into struct")
	}

	logger.Ctx(ctx).Trace().Msg("Config unmarshalled successfully")
}

// Read is a function that prepares the configuration data, reads it in from the
// config file, and then unmarshals it into the config struct. If any step fails,
// it logs a warning or fatal error as appropriate. If all steps succeed, it logs
// a trace message indicating success.
func Read(ctx context.Context) {
	preRead(ctx)

	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		switch {
		case errors.As(err, &viper.ConfigFileNotFoundError{}):
			logger.Ctx(ctx).Warn().Str("path", viper.ConfigFileUsed()).Msg("Config file not found. Skipping...")
		default:
			logger.Fatal().Err(err).Msg("Unable to read config file")
		}
	}

	ReadToStruct(ctx)
}

// preRead is a function that sets up the default configuration values and
// enables Viper's automatic environment variable handling. It logs a trace
// message indicating success.
func preRead(ctx context.Context) {
	for k, v := range _defaults {
		viper.SetDefault(k, v)
	}

	logger.Ctx(ctx).Trace().Msg("Config defaults set successfully")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	logger.Ctx(ctx).Trace().Msg("Config pre-read setup successfully")
}
