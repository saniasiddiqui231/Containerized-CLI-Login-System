package mfa

import "testing"

func TestGenerate(t *testing.T) {

	setup, err := Generate("testuser")

	if err != nil {
		t.Fatal(err)
	}

	if setup.Secret == "" {
		t.Fatal("secret should not be empty")
	}

	if setup.URL == "" {
		t.Fatal("url should not be empty")
	}
}
