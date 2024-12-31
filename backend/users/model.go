package users

import (
	"context"
	"time"

	"backend/configs"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID          string    `json:"id" bson:"_id" query:"id" form:"id" param:"id"`
	Email       string    `json:"email" bson:"email" query:"email" form:"email" param:"email"`
	Phone       string    `json:"phone" bson:"phone" query:"phone" form:"phone" param:"phone"`
	Password    []byte    `json:"password" bson:"password" query:"password" form:"password" param:"password"`
	Role        string    `json:"role" bson:"role" query:"role" form:"role" param:"role"`
	Status      string    `json:"status" bson:"status" query:"status" form:"status" param:"status"`
	AccountType string    `json:"account_type" bson:"account_type" query:"account_type" form:"account_type" param:"account_type"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at" query:"created_at" form:"created_at" param:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at" query:"updated_at" form:"updated_at" param:"updated_at"`
}

func (u *User) CreateNewUser() error {
	_, err := configs.StoreRequestInDb(*u, "users")
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(id string) User {
	var user User
	_ = configs.MI.DB.Collection("users").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return user
}

func GetUserByPhone(phone string) (User, error) {
	var user User
	err := configs.MI.DB.Collection("users").FindOne(context.TODO(), bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := configs.MI.DB.Collection("users").FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
