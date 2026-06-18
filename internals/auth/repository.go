package auth

import "CLI-login-system/internals/models"

type UserRepo interface {
	UsernameExists(username string) (bool, error)
	CreateUser(username, passwordHash string) error

	GetUserByUsername(username string) (*models.User, error)

	IncrementFailedAttempts(userID int64) error
	LockAccount(userID int64, minutes int) error
	ResetFailedAttempts(userID int64) error
}
