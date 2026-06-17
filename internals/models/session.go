package models

type Session struct {
	ID        int64
	UserID    int64
	Token     string
	ExpiresAt string
	Active    bool
}
