package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type (
	Config struct {
		App      App
		MasterDB DB
		Redis    Redis
	}

	App struct {
		Name string
		Env  string
		Port int
	}

	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		DB       string
	}

	Redis struct {
		Host     string
		Port     int
		Password string
		DB       int
	}
)

func New() *Config {
	return &Config{
		App: App{
			Name: getEnv("APP_NAME", "go-boilerplate"),
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnvAsInt("APP_PORT", 3000),
		},
		MasterDB: DB{
			Host:     getEnv("POSTGRES_HOST_MASTER", "localhost"),
			Port:     getEnvAsInt("POSTGRES_PORT_MASTER", 5432),
			User:     getEnv("POSTGRES_USER_MASTER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD_MASTER", "postgres"),
			DB:       getEnv("POSTGRES_DB_MASTER", "sagala_v1_db"),
		},
		Redis: Redis{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

// function getEnvAsInt is function to convert string to integer
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Database config
func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.MasterDB.User,
		c.MasterDB.Password,
		c.MasterDB.Host,
		c.MasterDB.Port,
		c.MasterDB.DB,
	)
}

// Redis config
func (c *Config) RedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	}
}
