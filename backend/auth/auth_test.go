package auth

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken(t *testing.T) {
	os.Setenv("SESSION_KEY", "test-secret-key-for-jwt-signing")
	defer os.Unsetenv("SESSION_KEY")

	secret := []byte(os.Getenv("SESSION_KEY"))

	claims := jwt.MapClaims{
		"name":    "Test User",
		"user_id": "test-user-123",
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	if tokenString == "" {
		t.Fatal("Expected non-empty token string")
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	if parsedClaims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		if parsedClaims["user_id"] != "test-user-123" {
			t.Errorf("user_id = %v, want %v", parsedClaims["user_id"], "test-user-123")
		}
		if parsedClaims["name"] != "Test User" {
			t.Errorf("name = %v, want %v", parsedClaims["name"], "Test User")
		}
	} else {
		t.Fatal("Failed to parse claims")
	}
}

func TestClaimsExpiration(t *testing.T) {
	claims := jwt.MapClaims{
		"name":    "User",
		"user_id": "123",
		"exp":     time.Now().Add(-1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err == nil {
		t.Error("Expected error for expired token, got nil")
	}
}

func TestInvalidSecret(t *testing.T) {
	claims := jwt.MapClaims{
		"user_id": "123",
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("correct-secret"))
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("wrong-secret"), nil
	})

	if err == nil {
		t.Error("Expected error for wrong secret, got nil")
	}
}
