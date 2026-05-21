package configs

import (
	"os"
	"testing"
)

func TestAppEnv(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	defer os.Unsetenv("APP_ENV")

	if got := AppEnv(); got != "test" {
		t.Errorf("AppEnv() = %q, want %q", got, "test")
	}
}

func TestEnvIsProd(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{"production", "production", true},
		{"development", "development", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("APP_ENV", tt.env)
			defer os.Unsetenv("APP_ENV")

			if got := EnvIsProd(); got != tt.want {
				t.Errorf("EnvIsProd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvPort(t *testing.T) {
	os.Setenv("PORT", "8080")
	defer os.Unsetenv("PORT")

	if got := EnvPort(); got != "8080" {
		t.Errorf("EnvPort() = %q, want %q", got, "8080")
	}
}

func TestEnvMongoURI(t *testing.T) {
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")
	defer os.Unsetenv("MONGODB_URI")

	if got := EnvMongoURI(); got != "mongodb://localhost:27017" {
		t.Errorf("EnvMongoURI() = %q, want %q", got, "mongodb://localhost:27017")
	}
}

func TestGetSessionKey(t *testing.T) {
	os.Setenv("SESSION_KEY", "secret-key")
	defer os.Unsetenv("SESSION_KEY")

	if got := GetSessionKey(); got != "secret-key" {
		t.Errorf("GetSessionKey() = %q, want %q", got, "secret-key")
	}
}

func TestEnvVarsReturnEmptyWhenUnset(t *testing.T) {
	vars := []struct {
		name string
		fn   func() string
	}{
		{"AppEnv", AppEnv},
		{"GetRedisUrl", GetRedisUrl},
		{"GetSessionKey", GetSessionKey},
		{"EnvMongoURI", EnvMongoURI},
		{"EnvPort", EnvPort},
		{"GetPaypackSecret", GetPaypackSecret},
		{"GetPaypackId", GetPaypackId},
		{"GetPlunkKey", GetPlunkKey},
		{"TelegramBotId", TelegramBotId},
		{"TelegramChatID", TelegramChatID},
	}

	for _, v := range vars {
		t.Run(v.name, func(t *testing.T) {
			os.Unsetenv("APP_ENV")
			os.Unsetenv("REDIS_URL")
			os.Unsetenv("SESSION_KEY")
			os.Unsetenv("MONGODB_URI")
			os.Unsetenv("PORT")
			os.Unsetenv("PAYPACK_CLIENT_SECRET")
			os.Unsetenv("PAYPACK_CLIENT_ID")
			os.Unsetenv("USE_PLUNK")
			os.Unsetenv("TELEGRAM_BOT_ID")
			os.Unsetenv("TELEGRAM_CHAT_ID")

			if got := v.fn(); got != "" {
				t.Errorf("%s() = %q, want empty string when env var is unset", v.name, got)
			}
		})
	}
}
