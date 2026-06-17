package database

import (
	"database/sql"
	"fmt"

	"CLI-login-system/internals/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) UsernameExists(username string) (bool, error) {
	var count int

	err := r.DB.QueryRow(
		"SELECT COUNT(*) FROM users WHERE username = ?",
		username,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepository) CreateUser(username, passwordHash string) error {
	_, err := r.DB.Exec(
		"INSERT INTO users(username, password_hash) VALUES(?, ?)",
		username,
		passwordHash,
	)

	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	err := r.DB.QueryRow(`
	SELECT
		id,
		username,
		password_hash,
		mfa_enabled,
		totp_secret,
		created_at,
		last_login,
		failed_attempts,
		locked_until
	FROM users
	WHERE username = ?
`, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.MFAEnabled,
		&user.TOTPSecret,
		&user.CreatedAt,
		&user.LastLogin,
		&user.FailedAttempts,
		&user.LockedUntil,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *UserRepository) UpdateLastLogin(userID int64) error {
	_, err := r.DB.Exec(`
		UPDATE users
		SET last_login = CURRENT_TIMESTAMP
		WHERE id = ?
	`, userID)

	return err
}
func (r *UserRepository) IncrementFailedAttempts(userID int64) error {
	_, err := r.DB.Exec(`
		UPDATE users
		SET failed_attempts = failed_attempts + 1
		WHERE id = ?
	`, userID)

	return err
}
func (r *UserRepository) LockAccount(userID int64, minutes int) error {
	_, err := r.DB.Exec(`
		UPDATE users
		SET locked_until = datetime('now', ?)
		WHERE id = ?
	`, fmt.Sprintf("+%d minutes", minutes), userID)

	return err
}
func (r *UserRepository) ResetFailedAttempts(userID int64) error {
	_, err := r.DB.Exec(`
		UPDATE users
		SET
			failed_attempts = 0,
			locked_until = NULL
		WHERE id = ?
	`, userID)

	return err
}
func (r *UserRepository) EnableMFA(
	userID int64,
	secret string,
) error {

	_, err := r.DB.Exec(`
		UPDATE users
		SET
			mfa_enabled = 1,
			totp_secret = ?
		WHERE id = ?
	`, secret, userID)

	return err
}
func (r *UserRepository) DisableMFA(
	userID int64,
) error {

	_, err := r.DB.Exec(`
		UPDATE users
		SET
			mfa_enabled = 0,
			totp_secret = NULL
		WHERE id = ?
	`, userID)

	return err
}
func (r *UserRepository) GetUserByID(id int64) (*models.User, error) {

	var user models.User

	err := r.DB.QueryRow(`
		SELECT
			id,
			username,
			password_hash,
			mfa_enabled,
			totp_secret,
			created_at,
			last_login,
			failed_attempts,
			locked_until
		FROM users
		WHERE id = ?
	`, id).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.MFAEnabled,
		&user.TOTPSecret,
		&user.CreatedAt,
		&user.LastLogin,
		&user.FailedAttempts,
		&user.LockedUntil,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
