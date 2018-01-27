package config

import (
	"os"
	"time"

	"github.com/talento90/gorpo/pkg/httpapi"
	"github.com/talento90/gorpo/pkg/log"
	"github.com/talento90/gorpo/pkg/repository/mongo"
	"github.com/talento90/gorpo/pkg/repository/redis"
)

// GetLogConfiguration get logger configurations
func GetLogConfiguration() (log.Configuration, error) {
	config := log.Configuration{
		Level:  "debug",
		Output: os.Stdout,
	}

	return config, config.Validate()
}

// GetServerConfiguration get server configurations
func GetServerConfiguration() (httpapi.Configuration, error) {
	config := httpapi.Configuration{
		Address:      ":4005",
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 4,
	}

	return config, config.Validate()
}

// GetMongoConfiguration returns the mongo configuration
func GetMongoConfiguration() (mongo.Configuration, error) {
	config := mongo.Configuration{
		Database: "gorpo",
		MongoURL: getEnv("MONGO_SERVICE_NAME", "localhost:27017"),
	}

	return config, config.Validate()
}

// GetRedisConfiguration returns the redis configuration
func GetRedisConfiguration() (redis.Configuration, error) {
	config := redis.Configuration{
		Address: getEnv("REDIS_SERVICE_NAME", "localhost:6379"),
	}

	return config, config.Validate()
}

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}