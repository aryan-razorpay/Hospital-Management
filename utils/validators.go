package utils

import (
	"errors"
	"regexp"
)

func ValidateContactNumber(contact string) error {
	if len(contact) != 10 {
		return errors.New("contact number must be 10 digits")
	}
	match, _ := regexp.MatchString(`^\d{10}$`, contact)
	if !match {
		return errors.New("invalid contact number format")
	}
	return nil
}

func ValidateIDLength(id string) error {
	if len(id) != 5 {
		return errors.New("ID must be 5 characters long")
	}
	return nil
}
func IsValidName(name string) bool {
	re := regexp.MustCompile("^[A-Za-z ]+$")
	return re.MatchString(name)
}
