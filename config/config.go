package config

type Config struct {
	// App configuration
	ListenHost              string  `envconfig:"LISTEN_HOST" default:"0.0.0.0"`
	GRPCPort                int     `envconfig:"GRPC_PORT" default:"9090"`
	HTTPPort                int     `envconfig:"HTTP_PORT" default:"9091"`
	AppName                 string  `envconfig:"APP_NAME" default:"ColdBrewService"`
	Environment             string  `envconfig:"ENVIRONMENT" default:""`
	LogLevel                string  `envconfig:"LOG_LEVEL" default:"info"`
	JSONLogs                bool    `envconfig:"JSON_LOGS" default:"true"`
	DisableSwagger          bool    `envconfig:"DISABLE_SWAGGER" default:"false"`
	DisableDebug            bool    `envconfig:"DISABLE_DEBUG" default:"false"`
	DisablePormetheus       bool    `envconfig:"DISABLE_PROMETHEUS" default:"false"`
	NewRelicLicenseKey      string  `envconfig:"NEW_RELIC_LICENSE_KEY" default:""`
	SentryDSN               string  `envconfig:"SENTRY_DSN" default:""`
	ReleaseName             string  `envconfig:"RELEASE_NAME" default:""`
	DisableGRPCReflection   bool    `envconfig:"DISABLE_GRPC_REFLECTION" default:"false"`
	JaegerLocalAgent        string  `envconfig:"JAEGER_LOCAL_AGENT" default:""`
	JaegerCollectorEndpoint string  `envconfig:"JAEGER_COLLECTOR_ENDPOINT" default:""`
	JaegerSamplerType       string  `envconfig:"JAEGER_SAMPLER_TYPE" default:""`
	JaegerSamplerParams     float64 `envconfig:"JAEGER_SAMPLER_PARAMS" default:""`
}
