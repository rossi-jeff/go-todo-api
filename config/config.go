package config

import "os"

type Config struct {
	Database DbConfig
	Secret   SecretConfig
}

func New() *Config {
	return &Config{
		Database: DbConfig{
			DbUser: getEnv("DB_USER", ""),
			DbPass: getEnv("DB_PASS", ""),
			DbHost: getEnv("DB_HOST", ""),
			DbName: getEnv("DB_NAME", ""),
			DbPort: getEnv("DB_PORT", ""),
		},
		Secret: SecretConfig{
			UserPass: getEnv("RAND_PW", "S3cr3t!"),
			JWT: getEnv("JWT_SECRET","S3cr3t!")
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
