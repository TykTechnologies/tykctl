package util

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
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

func SaveToConfig(key, token string) error {
	viper.Set(key, token)
	err := viper.WriteConfig()
	if err != nil {
		message := fmt.Sprintf("Couldn't write config: %s\n", err.Error())
		return errors.New(message)
	}
	return err

}
