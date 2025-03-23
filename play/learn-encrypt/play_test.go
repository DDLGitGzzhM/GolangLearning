package learn_encrypt

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func Test_bcrypt(t *testing.T) {
	password := "123456"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("encrypt pwd 1 : ", string(hashedPassword))

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("encrypt pwd 2 : ", string(hashedPassword))
}
