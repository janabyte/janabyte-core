package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func CheckPhoneNumber(phone string) error {
	// Trim any spaces from the phone number
	phone = strings.TrimSpace(phone)

	// Check if it starts with a plus (+) for international numbers
	if strings.HasPrefix(phone, "+") {
		// Strip the plus sign for further validation
		phone = phone[1:]
	}

	// Check the length of the phone number (between 10 and 15 digits)
	if len(phone) < 10 || len(phone) > 15 {
		return fmt.Errorf("Invalid length of the phone number")
	}

	//if !isDigitsOnly(phone) {
	//	return fmt.Errorf("Phone number must contain only digits")
	//}

	return nil
}

func isDigitsOnly(phone string) bool {
	// Regular expression to check if the phone string contains only digits
	re := regexp.MustCompile("^[0-9]+$")
	return re.MatchString(phone)
}

//87077765450
//+77077765450
