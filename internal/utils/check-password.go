package utils

import (
	"fmt"
	"regexp"
	"strings"
)

const MinPasswordLength = 12

func IsValidPassword(password string) error {

	if len(password) < MinPasswordLength {
		return fmt.Errorf("password must be at least %d characters", MinPasswordLength)
	}

	if !containsUppercase(password) {
		return fmt.Errorf("password must contain uppercase characters")
	}

	if !containsLowercase(password) {
		return fmt.Errorf("password must contain lowercase characters")
	}

	if !containsDigit(password) {
		return fmt.Errorf("password must contain digits")
	}

	if !containsSpecialCharacter(password) {
		return fmt.Errorf("password must contain special characters")
	}

	// Check for whitespace
	if containsWhitespace(password) {
		return fmt.Errorf("password must not contain any spaces")
	}

	return nil
}

func containsUppercase(password string) bool {
	re := regexp.MustCompile("[A-Z]")
	return re.MatchString(password)
}

func containsLowercase(password string) bool {
	re := regexp.MustCompile("[a-z]")
	return re.MatchString(password)
}

func containsDigit(password string) bool {
	re := regexp.MustCompile("[0-9]")
	return re.MatchString(password)
}

func containsSpecialCharacter(password string) bool {
	re := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	return re.MatchString(password)
}

func containsWhitespace(password string) bool {
	return strings.Contains(password, " ")
}
