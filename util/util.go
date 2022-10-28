package util

import (
	"errors"
	"net/mail"
	"strings"
)

// StringIsEmpty returns true is a string is empty.
func StringIsEmpty(v string) bool {

	return len(strings.TrimSpace(v)) == 0
}

// ValidateEmail validates that an email is of the right format.
func ValidateEmail(email string) error {
	if StringIsEmpty(email) {
		return errors.New("email cannot be empty")
	}
	_, err := mail.ParseAddress(email)
	return err
}
