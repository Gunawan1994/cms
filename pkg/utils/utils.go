package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	key               = []byte("secret")
	refreshTokenStore = make(map[string]string)
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateTokens(userID string, role string) (string, string, error) {
	claims := CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			Issuer:    "cms",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(key)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		Issuer:    "cms",
	})

	refreshTokenString, err := refreshToken.SignedString(key)
	if err != nil {
		return "", "", err
	}

	refreshTokenStore[refreshTokenString] = userID

	return accessTokenString, refreshTokenString, nil
}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	var claims CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}
