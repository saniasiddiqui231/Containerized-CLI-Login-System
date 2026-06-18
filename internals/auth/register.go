package auth

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo UserRepo
}

func (s *Service) Register(username, password string) error {

	username = strings.TrimSpace(username)

	if username == "" {
		return errors.New("username cannot be empty")
	}

	if password == "" {
		return errors.New("password cannot be empty")
	}

	exists, err := s.Repo.UsernameExists(username)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("username already exists")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	return s.Repo.CreateUser(
		username,
		string(hash),
	)
}
