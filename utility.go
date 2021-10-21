package main

import "net/mail"

// isEmail checks if the string is an email address
func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
