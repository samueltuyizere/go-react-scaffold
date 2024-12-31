package auth

import (
	"fmt"
	"time"

	"backend/users"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func processUserLogin(email, password string) (string, bool, error) {
	user, err := users.GetUserByEmail(email)
	if err != nil {
		return "", false, err
	}
	if user.Status != "active" {
		return "", false, fmt.Errorf("user is not active")
	}
	errr := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if errr != nil {
		return "", false, fmt.Errorf("the password provided is not valid")
	}

	return user.ID, true, nil

}

func createNewUser(email string, password []byte) (users.User, error) {
	uid := uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return users.User{}, err
	}
	user := users.User{
		Email:     email,
		Status:    "pending",
		Password:  hashedPassword,
		ID:        uid,
		CreatedAt: time.Now(),
	}
	err = user.CreateNewUser()
	if err != nil {
		return users.User{}, err
	}
	return user, nil
}
