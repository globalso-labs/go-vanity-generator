/*
 * go-vanity-generator
 * sourcehut.go
 * This file is part of go-vanity-generator.
 * Copyright (c) 2024.
 * Last modified at Mon, 24 Jul 2023 15:07:16 -0500 by nick.
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

package providers

import (
	"fmt"

	"go.globalso.dev/x/tools/vanity/config"
)

type sourcehutProvider struct {
	baseProvider
}

// GetGoSourceTag implements Provider.
func (p *sourcehutProvider) GetGoSourceTag() string {
	return fmt.Sprintf(
		"%s/%s %s %s/tree/%s/item{/dir} %s/tree/%s/item{/dir}/{file}#L{line}",
		p.domain, p.name,
		p.repo,
		p.repo, p.branch,
		p.repo, p.branch,
	)
}

var _ Provider = &gitlabProvider{}

func newSourcehutProvider(domain string, pkg config.Package) Provider {
	return &sourcehutProvider{
		baseProvider: baseProvider{
			domain: domain,
			name:   pkg.Name,
			repo:   pkg.URL,
			branch: pkg.Branch,
		},
	}
}
