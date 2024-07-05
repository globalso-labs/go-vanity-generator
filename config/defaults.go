package config

var _defaults = map[string]interface{}{
	"TELEMETRY.ENDPOINT":         "https://telemetry.idbi.pe/otlp",
	"TELEMETRY.DISABLED":         false,
	"TELEMETRY.METRICS_DISABLED": false,
	"TELEMETRY.TRACES_DISABLED":  false,
	"TELEMETRY.LOGS_DISABLED":    false,
}
