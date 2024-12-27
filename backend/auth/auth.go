package auth

import (
	"backend/configs"
	"backend/users"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

const (
	accessTokenCookieName = "Bearer"
)

func GetJWTSecret() string {
	return configs.GetSessionKey()
}

type Claims struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
	jwt.Claims
}

// GenerateTokensAndSetCookies generates jwt token and saves it to the http-only cookie.
func GenerateTokensAndSetCookies(user *users.User, c echo.Context) (string, error) {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return "", err
	}

	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	setUserCookie(user, exp, c)

	return accessToken, nil
}

func generateAccessToken(user *users.User) (string, time.Time, error) {
	// Declare the expiration time of the token (1h).
	expirationTime := time.Now().Add(24 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

// Pay attention to this function. It holds the main JWT token generation logic.
func generateToken(user *users.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time.
	claims := &Claims{
		UserId: user.ID,
		Claims: jwt.MapClaims{
			"exp": expirationTime.Unix(),
		},
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string.
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

// Here we are creating a new cookie, which will store the valid JWT token.
func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	// Http-only helps mitigate the risk of client side script accessing the protected cookie.
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

// Purpose of this cookie is to store the user's name.
func setUserCookie(user *users.User, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.ID
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

// JWTErrorChecker will be executed when user try to access a protected path.
func JWTErrorChecker(err error, c echo.Context) error {
	// Redirects to the signIn form.
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userSignInForm"))
}
