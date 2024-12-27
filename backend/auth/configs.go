package auth

import (
	"backend/configs"

	"github.com/gorilla/sessions"
)

var SessionStore = sessions.NewCookieStore([]byte(configs.GetSessionKey()))
