package validator

import (
	"net/mail"
	"strings"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsEmailFromProvider(email string, provider string) bool {
	at := strings.LastIndex(email, "@")
	if at == -1 {
		return false
	}
	domain := email[at+1:]

	return strings.Contains(domain, provider)
}
