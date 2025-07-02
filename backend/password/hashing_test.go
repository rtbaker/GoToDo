package password

import "testing"

func TestHashing(t *testing.T) {
	password := "password"

	hashedPassword, _ := HashPassword(password)

	if !VerifyPassword(password, hashedPassword) {
		t.Errorf("hashed password does not verify")
	}
}
