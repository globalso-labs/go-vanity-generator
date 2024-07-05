/*
 * go-template
 * main.go
 * This file is part of go-template.
 * Copyright (c) 2024.
 * Last modified at Sun, 24 Dec 2023 21:36:23 -0500 by nick.
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

package main

import (
	"github.com/nickaguilarh/credentials/cmd"

	"go.uber.org/automaxprocs/maxprocs"
)

// Then it builds the server and runs it.
func main() {
	// This controls the maxprocs environment variable in container runtimes.
	// https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	_, _ = maxprocs.Set()

	cmd.Execute()
}
