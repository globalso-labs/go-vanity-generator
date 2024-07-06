/*
 * go-vanity-generator
 * generator.go
 * This file is part of go-vanity-generator.
 * Copyright (c) 2024.
 * Last modified at Fri, 5 Jul 2024 20:51:07 -0500 by nick.
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

package generate

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gsols/go-logger"
	"go.globalso.dev/x/tools/vanity/config"
)

var defaultGenerator = new(DefaultGenerator)

type DefaultGenerator struct {
	// Output is the output directory for the generated files.
	output string

	// Clean determines whether we clean the output directory before generation.
	clean bool

	// Data is the data to be used in the templates.
	data map[string]interface{}
}

func (g *DefaultGenerator) SetClean(clean bool) {
	g.clean = clean
}

func (g *DefaultGenerator) SetOutput(output string) {
	path, err := filepath.Abs(output)
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get the absolute path. Using the original path")
		path = output
	}

	g.output = path
}

func (g *DefaultGenerator) Generate(ctx context.Context, vanity config.Vanity) error {
	ctx = logger.Ctx(ctx).With().
		Str("generator", "default").
		Str("domain", vanity.Domain).
		Logger().WithContext(ctx)

	// Clean the output directory if the flag is set.
	if g.clean {
		err := g.cleanPath(ctx)
		if err != nil {
			logger.Ctx(ctx).Error().Err(err).Msg("Failed to clean the output directory.")
			return err
		}
	}

	logger.Ctx(ctx).Info().Msgf("Generating vanity files for %d packages.", len(vanity.Packages))
	if err := g.assurePath(ctx); err != nil {
		logger.Ctx(ctx).Error().Err(err).Msg("Failed to assure the output directory.")
		return err
	}

	if err := g.generateStaticFiles(ctx, vanity); err != nil {
		logger.Ctx(ctx).Error().Err(err).Msg("Failed to generate the static files.")
		return err
	}

	if err := g.writePackagePages(ctx, vanity.Domain, vanity.Packages); err != nil {
		logger.Ctx(ctx).Error().Err(err).Msg("Failed to generate the package pages.")
		return err
	}

	return nil
}

// cleanPath removes all the files and directories in the output directory.
func (g *DefaultGenerator) cleanPath(ctx context.Context) error {
	logger.Ctx(ctx).Warn().Msgf("Cleaning the output directory %s.", g.output)

	if err := os.RemoveAll(g.output); err != nil {
		logger.Ctx(ctx).Trace().Str("path", g.output).Err(err).Msg("Failed to clean the output directory.")
		return err
	}

	return nil
}

// assurePath creates the output directory if it does not exist.
func (g *DefaultGenerator) assurePath(ctx context.Context) error {
	logger.Ctx(ctx).Trace().Msgf("Assuring the output directory %s.", g.output)

	if err := os.MkdirAll(g.output, os.ModePerm); err != nil {
		logger.Ctx(ctx).Trace().Str("path", g.output).Err(err).Msg("Failed to assure the output directory.")
		return err
	}

	return nil
}

func (g *DefaultGenerator) generateStaticFiles(ctx context.Context, params any) error {
	if err := g.writeErrorPage(ctx, params); err != nil {
		return err
	}

	if err := g.writeIndexPage(ctx, params); err != nil {
		return err
	}

	return nil
}
