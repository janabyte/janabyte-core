package utils

import (
	"fmt"
	"strings"
)

func CheckEmail(email string) error {
	if strings.Contains(email, "@") {
		return nil
	} else {
		return fmt.Errorf("Invalid email address: %s", email)
	}
}
