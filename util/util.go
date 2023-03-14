package util

import (
	"errors"
	"fmt"
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

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}

	return false
}

func AbbreviateDirection(direction string) string {
	switch direction {
	case "east":
		return "e"
	case "north":
		return "n"
	case "south":
		return "s"
	case "northeast":
		return "ne"
	case "northwest":
		return "nw"
	case "west":
		return "w"
	case "southwest":
		return "sw"
	case "southeast":
		return "se"
	case "central":
		return "c"
	}

	return ""
}

func GenerateURLFromZone(region string) (string, error) {
	regionPart := strings.Split(region, "-")

	if len(regionPart) != 4 {
		return "", errors.New("the format of this region is wrong")
	}

	return fmt.Sprintf("https://controller-aws-%s%s%s.cloud-ara.tyk.io:37001", regionPart[1], AbbreviateDirection(regionPart[2]), regionPart[3]), nil
}
