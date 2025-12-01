package token

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
	"intern/internal/configs"
	"intern/internal/core/domain"
	"intern/pkg/logger"
	"time"
)

func GenerateJWTToken(arg *domain.TokenRequest, log logger.ILogger) (*domain.TokenResponse, error) {
	accessNewJWt := jwt.New(jwt.SigningMethodHS256)

	accessClaims := accessNewJWt.Claims.(jwt.MapClaims)
	accessClaims["id"] = arg.ID
	accessClaims["role"] = arg.Role
	accessClaims["iat"] = time.Now().Unix()
	accessClaims["exp"] = time.Now().Add(configs.AccessExpireTime).Unix()

	accessToken, err := accessNewJWt.SignedString(configs.SignKey)
	if err != nil {
		log.Error("this error is signedString-~~~~~~~~~>ERROR", logger.Error(err))
		return nil, err
	}
	refreshNewJWT := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshNewJWT.Claims.(jwt.MapClaims)
	refreshClaims["id"] = arg.ID
	refreshClaims["role"] = arg.Role
	refreshClaims["iat"] = time.Now().Unix()
	refreshClaims["exp"] = time.Now().Add(configs.RefreshExpireTime).Unix()

	fmt.Println(configs.SignKey)

	refreshToken, err := refreshNewJWT.SignedString(configs.SignKey)
	if err != nil {
		log.Error("this error is signed  to string that sign key -~~~~~~~~~ERROR", logger.Error(err))
		return nil, err
	}

	return &domain.TokenResponse{
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		AccessExpiredTime:  cast.ToFloat64(accessClaims["exp"].(int64) - accessClaims["iat"].(int64)),
		RefreshExpiresTime: cast.ToFloat64(refreshClaims["exp"].(int64) - refreshClaims["iat"].(int64)),
		Success:            true,
	}, nil

}
