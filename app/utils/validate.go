package utils

import "regexp"

func IsEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}

func IsPhoneNumber(phoneNumber string) bool {
	pattern := `^(?:0|(\+){1,2}62)\d{6,11}\d{4}$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(phoneNumber)
}

func IsUri(uri string) bool {
	pattern := `^(([^:/?#]+):)(//([^/?#]*))([^?#]*)(\?([^#]*))?(#(.*))?`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(uri)
}
