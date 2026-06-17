package models

import "database/sql"

type User struct {
	ID           int64
	Username     string
	PasswordHash string

	MFAEnabled bool
	TOTPSecret sql.NullString

	CreatedAt string
	LastLogin sql.NullString

	FailedAttempts int
	LockedUntil    sql.NullString
}
