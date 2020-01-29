package authentication

import (
  "testing"
)

const (
  testPassword = "foobar"
)

// TestHashPassword hashes a password string and makes sure the result is not
// empty.
func TestHashPassword(t *testing.T) {
  hash, err := HashPassword(testPassword)
  if err != nil {
    t.Errorf("HashPassword returned err: %s", err.Error())
  }

  if len(hash) == 0 {
    t.Error("len(HashPassword) is 0")
  }
}

// TestIsPasswordCorrect hashes a password and then makes sure the comparison
// function returns true when the hash is compared against the original password
// string.
func TestIsPasswordCorrect(t *testing.T) {
  if hash, err := HashPassword(testPassword); err != nil {
    t.Errorf("HashPassword returned err: %s", err.Error())
  } else {
    if correct := IsPasswordCorrect(testPassword, hash); !correct {
      t.Error("IsPasswordCorrect returned false; expected true")
    }
  }
}
