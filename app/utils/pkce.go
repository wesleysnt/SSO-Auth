package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func GenerateCodeChallange(verifier string) (challenge string) {
	hash := sha256.Sum256([]byte(verifier))
	challenge = base64.StdEncoding.EncodeToString(hash[:])
	challenge = strings.Replace(challenge, "+", "-", -1)
	challenge = strings.Replace(challenge, "/", "_", -1)
	challenge = strings.Replace(challenge, "=", "", -1)
	return
}
