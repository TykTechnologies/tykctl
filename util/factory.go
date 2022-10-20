package util

import (
	"errors"
	"net/mail"
)

type Factory interface {
}

func StringIsEmpty(v string) bool {

	return len(v) == 0
}

func ValidateStringIsNotEmpty(value, message string) error {
	if StringIsEmpty(value) {
		return errors.New(message)
	}
	return nil
}

func ValidateEmail(email, message string) error {
	if len(email) == 0 {
		return errors.New(message)
	}
	_, err := mail.ParseAddress(email)
	return err
}

func ValidateNotEmpty(value string) error {
	if len(value) == 0 {
		return errors.New("cannot be empty string")
	}
	return nil
}
