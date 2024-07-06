/*
 * go-vanity-generator
 * static.go
 * This file is part of go-vanity-generator.
 * Copyright (c) 2024.
 * Last modified at Fri, 5 Jul 2024 22:13:36 -0500 by nick.
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
	"go.globalso.dev/x/tools/vanity/internal/data"
)

func (g *Gen) writeErrorPage(ctx context.Context, params any) error {
	var errorPath = filepath.Join(g.output, "404.html")

	tmpl, err := template.ParseFS(data.Templates, path.Join(data.TemplatesPath, data.ErrorPage))
	if err != nil {
		logger.Error().Err(err).Interface("templates", data.Templates).Msg("Failed to parse the error page template.")
		return err
	}

	logger.Ctx(ctx).Debug().Str("path", errorPath).Msg("Writing the error page.")
	file, err := os.Create(errorPath)
	if err != nil {
		logger.Error().Err(err).Str("path", errorPath).Msg("Failed to create the error page file.")
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, params)
}

func (g *Gen) writeIndexPage(ctx context.Context, params any) error {
	var indexPath = filepath.Join(g.output, "index.html")

	tmpl, err := template.ParseFS(data.Templates, path.Join(data.TemplatesPath, data.IndexPage))
	if err != nil {
		logger.Error().Err(err).Interface("templates", data.Templates).Msg("Failed to parse the index page template.")
		return err
	}

	logger.Ctx(ctx).Debug().Str("path", indexPath).Msg("Writing the index page.")
	file, err := os.Create(indexPath)
	if err != nil {
		logger.Error().Err(err).Str("path", indexPath).Msg("Failed to create the index page file.")
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, params)
}
