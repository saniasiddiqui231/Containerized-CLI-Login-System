package session

import "testing"

func TestGenerateToken(t *testing.T) {

	token1 := generateToken()
	token2 := generateToken()

	if token1 == "" {
		t.Fatal("token should not be empty")
	}

	if token2 == "" {
		t.Fatal("token should not be empty")
	}

	if token1 == token2 {
		t.Fatal("generated tokens should be unique")
	}
}
