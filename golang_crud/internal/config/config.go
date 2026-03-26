package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURL   string
	MongoDB    string
	ServerPort string
}

func Load() (Config, error) {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("unable to read env file")
	}

	mongoURL, err := extractEnv("MONGO_URL")
	if err != nil {
		return Config{}, err
	}

	mongoDB, err := extractEnv("MONGODB_NAME")
	if err != nil {
		return Config{}, err
	}

	port, err := extractEnv("MONGO_PORT_NO")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoURL:   mongoURL,
		MongoDB:    mongoDB,
		ServerPort: port,
	}, nil
}

func extractEnv(key string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("missing required env: %s", key)
	}

	return val, nil
}
