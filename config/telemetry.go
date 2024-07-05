/*
 * credentials
 * telemetry.go
 * This file is part of credentials.
 * Copyright (c) 2024.
 * Last modified at Fri, 5 Jul 2024 02:42:22 -0500 by nick.
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

// Telemetry holds the configuration for telemetry. We use OTLP as the unique
// exporter and all related configuration is set by environment variables.
type Telemetry struct {
	// GloballyDisabled is a flag to disable all telemetry.
	Disabled bool `mapstructure:"disabled"`

	// MetricsDisabled, TracesDisabled, and LogsDisabled are flags to disable
	// each type of telemetry.
	MetricsDisabled bool `mapstructure:"metrics_disabled"`
	TracesDisabled  bool `mapstructure:"traces_disabled"`
	LogsDisabled    bool `mapstructure:"logs_disabled"`

	// Endpoint is the endpoint to send telemetry data.
	Endpoint string `mapstructure:"endpoint"`
}
