package crypto

import (
	"testing"
	config2 "users_api/internal/pkg/config"
)

func TestHashAndSalt(t *testing.T) {
	password := []byte("password123")
	hashedPwd := HashAndSalt(password)

	if len(hashedPwd) == 0 {
		t.Errorf("Expected hashed password, but got empty string")
	}
}

func TestComparePasswords(t *testing.T) {
	password := []byte("password123")
	hashedPwd := HashAndSalt(password)

	match := ComparePasswords(hashedPwd, password)
	if !match {
		t.Errorf("Expected passwords to match, but got not matched")
	}

	wrongPassword := []byte("wrongpassword")
	match = ComparePasswords(hashedPwd, wrongPassword)
	if match {
		t.Errorf("Expected passwords not to match, but got matched")
	}
}

func TestCreateToken(t *testing.T) {
	config2.Config = &config2.Configuration{}
	config2.Config.Server = config2.ServerConfiguration{Secret: "secret"}
	username := "testuser"
	token, err := CreateToken(username)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(token) == 0 {
		t.Errorf("Expected token, but got empty string")
	}
}

func TestValidateToken(t *testing.T) {
	config2.Config = &config2.Configuration{}
	config2.Config.Server = config2.ServerConfiguration{Secret: "secret"}

	username := "testuser"
	token, _ := CreateToken(username)

	mockToken := "Bearer " + token

	valid := ValidateToken(mockToken)

	if !valid {
		t.Errorf("Expected token to be valid, but got invalid")
	}
}
