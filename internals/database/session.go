package database

import (
	"database/sql"
	"time"
)

type SessionRepository struct {
	DB *sql.DB
}

func (r *SessionRepository) CreateSession(
	userID int64,
	token string,
	expiry time.Time,
) error {

	_, err := r.DB.Exec(`
		INSERT INTO sessions(
			user_id,
			token,
			expires_at,
			active
		)
		VALUES(?, ?, ?, 1)
	`,
		userID,
		token,
		expiry.Format("2006-01-02 15:04:05"),
	)

	return err
}

func (r *SessionRepository) DeactivateSession(
	token string,
) error {

	_, err := r.DB.Exec(`
		UPDATE sessions
		SET active = 0
		WHERE token = ?
	`, token)

	return err
}
