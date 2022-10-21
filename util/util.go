package util

import (
	"errors"
	"net/mail"
	"strings"
)

func StringIsEmpty(v string) bool {

	return len(strings.TrimSpace(v)) == 0
}

func ValidateEmail(email string) error {
	if StringIsEmpty(email) {
		return errors.New("email cannot be empty")
	}
	_, err := mail.ParseAddress(email)
	return err
}
