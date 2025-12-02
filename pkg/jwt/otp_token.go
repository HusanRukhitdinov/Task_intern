package token

import (
	"crypto/rand"
	"errors"
	"fmt"
	"intern/internal/configs"
	"intern/internal/core/domain"

	"github.com/golang-jwt/jwt"
)

func DecodeOTPToken(tokenStr string) (*domain.OTPTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.OTPTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.SignKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*domain.OTPTokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid otp token")
}

func Generate6DigitCode() (string, error) {
	b := make([]byte, 3)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	num := int(b[0])<<16 | int(b[1])<<8 | int(b[2])
	code := num % 1000000
	return fmt.Sprintf("%06d", code), nil
}
