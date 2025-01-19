package middleware

import "testing"

func TestJWT(t *testing.T) {
	tokena, err := GenerateJWT("masha")
	if err != nil {
		t.Log("Failed to generate")
		t.Fail()
	}
	tokenb, err := GenerateJWT("Masha")
	if err != nil {
		t.Log("Failed to generate")
		t.Fail()
	}
	if tokena == tokenb {
		t.Log("Same token")
		t.Fail()
	}
	if err = ValidateJWT(tokena); err != nil {
		t.Log("Failed to validate")
		t.Fail()
	}
	if err = ValidateJWT(tokenb); err != nil {
		t.Log("Failed to validate")
		t.Fail()
	}
	if err = ValidateJWT("invalid-token"); err == nil {
		t.Log("Invalid token passed validation")
		t.Fail()
	}
}
