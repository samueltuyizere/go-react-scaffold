package utils

import (
	"backend/integrations"
	"fmt"
)

func SendOtpVerificationEmail(code, email, userId string) error {
	payload := fmt.Sprintf("Hello, \n\nYour OTP code is %s. \n\nPlease use this code to verify your email address. \n\nThank you.", code)
	err := integrations.SendEmailWithPlunk(payload, email, "Welcome to My Project", "mail@example.com")
	return err
}
