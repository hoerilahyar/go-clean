package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *service) GenerateAccessToken(
	userID uint64,
	username string,
) (string, error) {

	now := time.Now()

	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   fmt.Sprintf("%d", userID),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTokenTTL)),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.secretKey)
}

func (s *service) ParseAccessToken(
	tokenString string,
) (*Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return s.secretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (s *service) ValidateAccessToken(
	tokenString string,
) (*Claims, error) {

	claims, err := s.ParseAccessToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt == nil {
		return nil, fmt.Errorf("token has no expiration")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
