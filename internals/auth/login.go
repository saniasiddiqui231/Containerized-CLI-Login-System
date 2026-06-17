package auth

import (
	"CLI-login-system/internals/config"
	"CLI-login-system/internals/models"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(username, password string) (*models.User, error) {

	user, err := s.Repo.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user.LockedUntil.Valid {

		lockTime, err := time.Parse(
			time.RFC3339,
			user.LockedUntil.String,
		)

		if err != nil {
			lockTime, err = time.Parse(
				"2006-01-02 15:04:05",
				user.LockedUntil.String,
			)
		}

		if err == nil && lockTime.After(time.Now()) {
			return nil, errors.New(
				"account temporarily locked. try again later",
			)
		}
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {

		err = s.Repo.IncrementFailedAttempts(user.ID)
		if err != nil {
			return nil, err
		}

		if user.FailedAttempts+1 >= config.MaxFailedAttempts {

			err = s.Repo.LockAccount(
				user.ID,
				config.LockoutMinutes,
			)
			if err != nil {
				return nil, err
			}

			return nil, errors.New(
				"account temporarily locked. try again later",
			)
		}

		return nil, errors.New("invalid credentials")
	}

	err = s.Repo.ResetFailedAttempts(user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
