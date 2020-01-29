package auth

import (
  "testing"
  "github.com/dgrijalva/jwt-go"
)

const (
  testSecretKey = "testing"
)

// TestNewJWT tests the NewJWT function and ensures a token is returned when a
// specific set of custom claims are passed in.
func TestNewJWT(t *testing.T) {
  tkn := NewJWT(
    &jwt.StandardClaims{
      Audience: "tests",
      Subject: "1",
    },
  )

  if tkn == nil {
    t.Errorf("NewJWT returned nil; expected *jwt.Token")
  }
}

// IsJWTStrValid tests the IsJWTStrValid valid method and ensures a valid token
// can be 1. generated and 2. be tested to be valid.
func TestIsJWTStrValid(t *testing.T) {
  tkn := NewJWT(
    &jwt.StandardClaims{
      Audience: "tests",
      Subject: "1",
    },
  )

  tknStr, err := tkn.SignedString([]byte(testSecretKey))
  if err != nil {
    t.Errorf("tkn.SignedString returned error: %s", err.Error())
  }

  if valid, err := IsJWTStrValid(tknStr, testSecretKey); err != nil {
    t.Errorf("IsJWTStrValid returned error: %s", err.Error())
  } else if !valid {
    t.Error("IsJWTStrValid returned invalid (false); expected valid (true)")
  }
}

// ...
func TestParseJWTStr(t *testing.T) {

  // ...

}
