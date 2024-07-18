package facades

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sso-auth/app/utils/otel"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var HmacSecret []byte

type CustomClaim struct {
	UserId   uint   `json:"user_id"`
	ClientId string `json:"client_id"`
	jwt.RegisteredClaims
}

func ValidMAC(key []byte) {
	mac := hmac.New(sha256.New, key)
	HmacSecret = mac.Sum(nil)
}

func GenerateToken(secret, clientId string, userId, expiredDuration uint, ctx context.Context) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	_, span := otel.StartNewTrace(ctx, "GenerateToken")
	defer otel.EndSpan(span)

	if secret == "" {
		secret = os.Getenv("JWT_SECRET")
	}
	claim := CustomClaim{
		userId,
		clientId,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiredDuration) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

func GenerateSecret(s string) []byte {
	ValidMAC([]byte(s))
	return HmacSecret
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func ParseToken(tokenString, secret string, ctx context.Context) (token *jwt.Token, err error) {
	_, span := otel.StartNewTrace(ctx, "ParseToken")
	defer otel.EndSpan(span)
	if secret == "" {
		secret = os.Getenv("JWT_SECRET")
	}

	token, err = jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*CustomClaim); ok {
		fmt.Println(claims.UserId, claims.RegisteredClaims.Issuer)
	} else {
		return nil, errors.New("unknown claims type, cannot proceed")
	}

	return

}

func GetUserIdFromToken(tokenString, secret string, ctx context.Context) (uint, error) {
	token, err := ParseToken(tokenString, secret, ctx)
	if err != nil {
		return 0, err
	}

	claims := token.Claims.(*CustomClaim)
	return claims.UserId, nil
}

func GetClientIdFromToken(tokenString, secret string, ctx context.Context) (string, error) {
	token, err := ParseToken(tokenString, secret, ctx)
	if err != nil {
		return "", err
	}

	claims := token.Claims.(*CustomClaim)
	return claims.ClientId, nil
}
