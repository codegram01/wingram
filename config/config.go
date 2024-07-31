package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	DbHost, DbPort, DbUser, DbName, DbPass, DbSSL string
}

func Init() *Config {
	cfg := &Config{
		Port: GetEnv("PORT", "8080"),

		DbHost: GetEnv("DB_HOST", "localhost"),
		DbPort: GetEnv("DB_PORT", "5432"),
		DbUser: GetEnv("DB_USER", "testuser"),
		DbName: GetEnv("DB_NAME", "testdb"),
		DbPass: GetEnv("DB_PASS", "testpassword"),
		DbSSL : GetEnv("DB_SSL", "disable"),
	}
	return cfg
}

// this function load env value
// if not have value it will assign defaultVal
func GetEnv(key string, defaultVal string) string {
	val := os.Getenv(key)

	if val == "" {
		val = defaultVal
	}

	return val
}

func (cfg *Config) GetDbConStr() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", 
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbName,
		cfg.DbPass,
		cfg.DbSSL,
	)
}