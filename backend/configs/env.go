package configs

import (
	"os"
)

func AppEnv() string {
	return os.Getenv("APP_ENV")
}

func EnvIsProd() bool {
	return AppEnv() == "production"
}


func GetRedisUrl() string {
	return os.Getenv("REDIS_URL")
}

func GetSessionKey() string {
	return os.Getenv("SESSION_KEY")
}

func EnvMongoURI() string {
	return os.Getenv("MONGODB_URI")
}

func EnvPort() string {
	return os.Getenv("PORT")
}
