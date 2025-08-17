package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Config holds application settings
	type Config struct {
		PostgresHost     string `json:"postgres_host"`
		PostgresPort     int    `json:"postgres_port"`
		PostgresUser     string `json:"postgres_user"`
		PostgresPassword string `json:"postgres_password"`
		PostgresDatabase string `json:"postgres_database"`
		SQLiteDatabase   string `json:"sqlite_database"`
		RedisHost        string `json:"redis_host"`
		RedisPort        int    `json:"redis_port"`
		RedisPassword    string `json:"redis_password"`
		CacheTTL         int    `json:"cache_ttl"`
		PrismaBinary     string `json:"prisma_binary"`
	}

// LoadConfig loads application settings from environment variables and command-line arguments
func LoadConfig() (*Config, error) {
	var config Config

	// Load from environment variables
	config.PostgresHost = getEnv("POSTGRES_HOST", "localhost")
	config.PostgresPort, _ = strconv.Atoi(getEnv("POSTGRES_PORT", "5432"))
	config.PostgresUser = getEnv("POSTGRES_USER", "postgres")
	config.PostgresPassword = getEnv("POSTGRES_PASSWORD", "postgres")
	config.PostgresDatabase = getEnv("POSTGRES_DATABASE", "data-sync")
	config.SQLiteDatabase = getEnv("SQLITE_DATABASE", "data-sync.db")
	config.RedisHost = getEnv("REDIS_HOST", "localhost")
	config.RedisPort, _ = strconv.Atoi(getEnv("REDIS_PORT", "6379"))
	config.RedisPassword = getEnv("REDIS_PASSWORD", "")
	config.CacheTTL, _ = strconv.Atoi(getEnv("CACHE_TTL", "3600"))
	config.PrismaBinary = getEnv("PRISMA_BINARY", "npx prisma")

	// Load from command-line arguments
	flag.StringVar(&config.PostgresHost, "postgres-host", config.PostgresHost, "Postgres host")
	flag.IntVar(&config.PostgresPort, "postgres-port", config.PostgresPort, "Postgres port")
	flag.StringVar(&config.PostgresUser, "postgres-user", config.PostgresUser, "Postgres user")
	flag.StringVar(&config.PostgresPassword, "postgres-password", config.PostgresPassword, "Postgres password")
	flag.StringVar(&config.PostgresDatabase, "postgres-database", config.PostgresDatabase, "Postgres database")
	flag.StringVar(&config.SQLiteDatabase, "sqlite-database", config.SQLiteDatabase, "SQLite database")
	flag.StringVar(&config.RedisHost, "redis-host", config.RedisHost, "Redis host")
	flag.IntVar(&config.RedisPort, "redis-port", config.RedisPort, "Redis port")
	flag.StringVar(&config.RedisPassword, "redis-password", config.RedisPassword, "Redis password")
	flag.IntVar(&config.CacheTTL, "cache-ttl", config.CacheTTL, "Cache TTL")
	flag.StringVar(&config.PrismaBinary, "prisma-binary", config.PrismaBinary, "Prisma binary")
	flag.Parse()

	return &config, nil
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

func (c *Config) Validate() error {
	if c.PostgresHost == "" {
		return fmt.Errorf("postgres host is required")
	}
	if c.PostgresPort == 0 {
		return fmt.Errorf("postgres port is required")
	}
	if c.PostgresUser == "" {
		return fmt.Errorf("postgres user is required")
	}
	if c.PostgresPassword == "" {
		return fmt.Errorf("postgres password is required")
	}
	if c.PostgresDatabase == "" {
		return fmt.Errorf("postgres database is required")
	}
	if c.SQLiteDatabase == "" {
		return fmt.Errorf("sqlite database is required")
	}
	if c.RedisHost == "" {
		return fmt.Errorf("redis host is required")
	}
	if c.RedisPort == 0 {
		return fmt.Errorf("redis port is required")
	}
	if c.CacheTTL == 0 {
		return fmt.Errorf("cache ttl is required")
	}
	if c.PrismaBinary == "" {
		return fmt.Errorf("prisma binary is required")
	}
	return nil
}

func (c *Config) String() string {
	json, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
