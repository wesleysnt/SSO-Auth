package utils

import (
	"context"
	"sso-auth/app/utils/otel"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, ctx context.Context) (string, error) {
	_, span := otel.StartNewTrace(ctx, "HashPassword")
	defer otel.EndSpan(span)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string, ctx context.Context) bool {
	_, span := otel.StartNewTrace(ctx, "CheckPasswordHash")
	defer otel.EndSpan(span)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
