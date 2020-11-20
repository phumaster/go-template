package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	expect := "123"
	password := "123"
	passwordHashed, err := HashPassword(password)
	if err != nil {
		t.Error("Error when hash password")
	}
	if !CheckPasswordHash(expect, passwordHashed) {
		t.Errorf("Check password hashed fail, expect %s. Reality %v", expect, passwordHashed)
	}
}
