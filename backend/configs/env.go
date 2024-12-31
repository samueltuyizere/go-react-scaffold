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

func GetPaypackSecret() string {
	return os.Getenv("PAYPACK_CLIENT_SECRET")
}

func GetPaypackId() string {
	return os.Getenv("PAYPACK_CLIENT_ID")
}

func GetPlunkKey() string {
	return os.Getenv("USE_PLUNK")
}

func TelegramBotId() string {
	return os.Getenv("TELEGRAM_BOT_ID")
}

func TelegramChatID() string {
	return os.Getenv("TELEGRAM_CHAT_ID")
}
