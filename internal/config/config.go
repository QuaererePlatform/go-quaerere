package config

type (
	DBInitConfig struct {
		AppEnv         *string `mapstructure:"app_env"`
		DebugMode      bool    `mapstructure:"debug_mode"`
		StorageBackend string  `mapstructure:"storage_backend"`
	}
)

