/*
 * go-vanity-generator
 * generate.go
 * This file is part of go-vanity-generator.
 * Copyright (c) 2024.
 * Last modified at Fri, 5 Jul 2024 20:14:33 -0500 by nick.
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

package generate

import (
	"context"
	"runtime"

	"github.com/gsols/go-logger"
	"github.com/spf13/viper"
	"go.globalso.dev/x/tools/vanity/config"
	"go.globalso.dev/x/tools/vanity/internal"
	"go.globalso.dev/x/tools/vanity/internal/constants/cmd"
)

// Execute runs the application and handles the shutdown signal.
func Execute(ctx context.Context) error {
	ctx = logger.Ctx(ctx).With().Str("module", "generate").Logger().WithContext(ctx)

	logger.Ctx(ctx).Info().Msgf("Running version %s on %s/%s with %d CPUs and %d GOMAXPROCS.",
		internal.Version, runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), runtime.GOMAXPROCS(0),
	)

	// Set the clean flag for the generator.
	defaultGenerator.SetClean(viper.GetBool(cmd.GeneratorClean))

	// Set the output directory for the generator.
	defaultGenerator.SetOutput(viper.GetString(cmd.GeneratorOutput))

	// Generate the vanity files.
	return defaultGenerator.Generate(ctx, config.Get().Vanity)
}
