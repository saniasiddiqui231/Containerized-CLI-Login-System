package session

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"CLI-login-system/internals/config"
	"CLI-login-system/internals/database"
)

type Service struct {
	Repo *database.SessionRepository
}

func generateToken() string {

	b := make([]byte, 32)

	_, _ = rand.Read(b)

	return hex.EncodeToString(b)
}

func (s *Service) Create(
	userID int64,
) (string, time.Time, error) {

	token := generateToken()

	expiry := time.Now().
		Add(time.Minute * config.SessionTimeoutMinutes)

	err := s.Repo.CreateSession(
		userID,
		token,
		expiry,
	)

	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiry, nil
}
