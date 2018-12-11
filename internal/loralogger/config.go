package loralogger

// Config holds the lora-logger config.
type Config struct {
	Bind     string          `mapstructure:"bind"`
	Backends []BackendConfig `mapstructure:"backend"`
}

// BackendConfig holds the config for a single backend.
type BackendConfig struct {
	Host       string   `mapstructure:"host"`
	GatewayIDs []string `mapstructure:"gateway_ids"`
}
