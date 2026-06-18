package config

import "testing"

func TestConfigurationValues(t *testing.T) {

	if MaxFailedAttempts <= 0 {
		t.Fatal("MaxFailedAttempts must be positive")
	}

	if LockoutMinutes <= 0 {
		t.Fatal("LockoutMinutes must be positive")
	}

	if SessionTimeoutMinutes <= 0 {
		t.Fatal("SessionTimeoutMinutes must be positive")
	}
}
