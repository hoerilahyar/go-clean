package jwt

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

func (s *service) GenerateRefreshToken() (string, time.Time, error) {

	// 32 bytes = 64 karakter hex
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", time.Time{}, err
	}

	refreshToken := hex.EncodeToString(b)

	expiredAt := time.Now().Add(s.refreshTokenTTL)

	return refreshToken, expiredAt, nil
}
