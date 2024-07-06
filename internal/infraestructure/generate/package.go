/*
 * go-vanity-generator
 * package.go
 * This file is part of go-vanity-generator.
 * Copyright (c) 2024.
 * Last modified at Fri, 5 Jul 2024 22:56:49 -0500 by nick.
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
	"html/template"
	"os"
	"path"
	"path/filepath"

	"github.com/gsols/go-logger"
	"go.globalso.dev/x/tools/vanity/config"
	"go.globalso.dev/x/tools/vanity/internal/data"
)

func (g *Gen) writePackagePages(ctx context.Context, domain string, packages []config.Package) error {
	tmpl, err := template.ParseFS(data.Templates, path.Join(data.TemplatesPath, data.PackagePage))
	if err != nil {
		logger.Ctx(ctx).Error().Err(err).Msg("Failed to parse the package page template.")
		return err
	}

	logger.Ctx(ctx).Debug().Str("path", g.output).Msg("Writing the package pages.")
	for _, p := range packages {
		if err = g.writePackage(ctx, tmpl, domain, p); err != nil {
			logger.Ctx(ctx).Error().Err(err).Str("package", p.Name).Msg("Failed to write the package page.")
			return err
		}
	}

	return nil
}

func (g *Gen) writePackage(ctx context.Context, tpl *template.Template, domain string, pkg config.Package) error {
	ctx = logger.Ctx(ctx).With().Str("package", pkg.Name).Logger().WithContext(ctx)

	var packagePath = filepath.Join(g.output, pkg.Name, "index.html")

	err := os.MkdirAll(filepath.Dir(packagePath), os.ModePerm)
	if err != nil {
		logger.Ctx(ctx).Error().Err(err).Str("path", packagePath).Msg("Failed to create the package page directory")
		return err
	}

	logger.Ctx(ctx).Trace().Msg("Writing the package pages.")
	file, err := os.Create(packagePath)
	if err != nil {
		logger.Ctx(ctx).Error().Err(err).Str("path", packagePath).Msg("Failed to create the package page file")
		return err
	}
	defer file.Close()

	attributes := NewPackageAttributes(ctx, domain, pkg)
	if err = tpl.Execute(file, attributes); err != nil {
		logger.Ctx(ctx).Error().Err(err).Msg("Failed to write the package page")
		return err
	}

	logger.Ctx(ctx).Info().Msg("Package page written successfully")
	return g.writeSubPackages(ctx, tpl, domain, pkg)
}

func (g *Gen) writeSubPackages(ctx context.Context, tpl *template.Template, domain string, pkg config.Package) error {
	if len(pkg.Subpackages) == 0 {
		return nil
	}

	logger.Ctx(ctx).Debug().Msgf("Writing %d subpackages fr", len(pkg.Subpackages))
	for _, p := range pkg.Subpackages {
		ctx = logger.Ctx(ctx).With().Str("subpackage", p).Logger().WithContext(ctx)

		pkg.Name = path.Join(pkg.Name, p)
		pkg.Subpackages = nil

		logger.Ctx(ctx).Trace().Msg("Writing the subpackage page")
		if err := g.writePackage(ctx, tpl, domain, pkg); err != nil {
			logger.Ctx(ctx).Error().Err(err).Msg("Failed to write the subpackage page")
			return err
		}
	}

	return nil
}
