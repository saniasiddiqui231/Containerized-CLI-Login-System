package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHashMatches(t *testing.T) {

	password := "secret123"

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword(
		hash,
		[]byte(password),
	)

	if err != nil {
		t.Fatal("password should match")
	}
}

func TestWrongPasswordFails(t *testing.T) {

	password := "secret123"

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword(
		hash,
		[]byte("wrong-password"),
	)

	if err == nil {
		t.Fatal("wrong password should fail")
	}
}
