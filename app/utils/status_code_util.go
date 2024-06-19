package utils

import "sso-auth/app/schemas"

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code schemas.ApiStatusCode) string {
	if text, ok := schemas.StatusTextMap[code]; ok {
		return text
	}

	return ""
}
