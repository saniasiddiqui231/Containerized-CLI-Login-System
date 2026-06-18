package auth

import (
	"CLI-login-system/internals/config"
	"CLI-login-system/internals/models"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type MockRepo struct {
	User   *models.User
	Exists bool

	IncrementCalled bool
	ResetCalled     bool
	LockCalled      bool
}

func (m *MockRepo) UsernameExists(username string) (bool, error) {
	return m.Exists, nil
}

func (m *MockRepo) CreateUser(username, hash string) error {
	return nil
}

func (m *MockRepo) GetUserByUsername(username string) (*models.User, error) {
	return m.User, nil
}

func (m *MockRepo) IncrementFailedAttempts(userID int64) error {
	m.IncrementCalled = true
	return nil
}

func (m *MockRepo) LockAccount(userID int64, minutes int) error {
	m.LockCalled = true
	return nil
}

func (m *MockRepo) ResetFailedAttempts(userID int64) error {
	m.ResetCalled = true
	return nil
}
func TestDuplicateUsername(t *testing.T) {

	repo := &MockRepo{
		Exists: true,
	}

	service := &Service{
		Repo: repo,
	}

	err := service.Register(
		"john",
		"password123",
	)

	if err == nil {
		t.Fatal("expected duplicate username error")
	}
}
func TestRegisterSuccess(t *testing.T) {

	repo := &MockRepo{
		Exists: false,
	}

	service := &Service{
		Repo: repo,
	}

	err := service.Register(
		"john",
		"password123",
	)

	if err != nil {
		t.Fatal(err)
	}
}
func TestLoginSuccess(t *testing.T) {

	hash, _ := bcrypt.GenerateFromPassword(
		[]byte("secret123"),
		bcrypt.DefaultCost,
	)

	repo := &MockRepo{
		User: &models.User{
			ID:           1,
			PasswordHash: string(hash),
		},
	}

	service := &Service{
		Repo: repo,
	}

	_, err := service.Login(
		"john",
		"secret123",
	)

	if err != nil {
		t.Fatal(err)
	}

	if !repo.ResetCalled {
		t.Fatal("expected failed attempts reset")
	}
}
func TestLoginWrongPassword(t *testing.T) {

	hash, _ := bcrypt.GenerateFromPassword(
		[]byte("secret123"),
		bcrypt.DefaultCost,
	)

	repo := &MockRepo{
		User: &models.User{
			ID:           1,
			PasswordHash: string(hash),
		},
	}

	service := &Service{
		Repo: repo,
	}

	_, err := service.Login(
		"john",
		"wrong-password",
	)

	if err == nil {
		t.Fatal("expected error")
	}

	if !repo.IncrementCalled {
		t.Fatal("expected failed attempt increment")
	}
}
func TestLockoutTriggered(t *testing.T) {

	hash, _ := bcrypt.GenerateFromPassword(
		[]byte("secret123"),
		bcrypt.DefaultCost,
	)

	repo := &MockRepo{
		User: &models.User{
			ID:             1,
			PasswordHash:   string(hash),
			FailedAttempts: config.MaxFailedAttempts - 1,
		},
	}

	service := &Service{
		Repo: repo,
	}

	_, err := service.Login(
		"john",
		"wrong-password",
	)

	if err == nil {
		t.Fatal("expected lockout")
	}

	if !repo.LockCalled {
		t.Fatal("expected lock account")
	}
}
