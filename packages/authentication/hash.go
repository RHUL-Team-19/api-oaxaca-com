package authentication

import (
  "golang.org/x/crypto/bcrypt"
)

// HashPassword takes a password string and returns a bcrypt hash of the
// password.
func HashPassword(password string) ([]byte, error) {
  h, err := bcrypt.GenerateFromPassword([]byte(password), 11)
  if err != nil {
    return nil, err
  } else {
    return h, nil
  }
}

// IsPasswordCorrect takes a password string and a hash byte slice and returns
// true if the password matches, else false.
func IsPasswordCorrect(password string, hash []byte) bool {
  err := bcrypt.CompareHashAndPassword(hash, []byte(password))
  if err == nil {
    return true
  } else {
    return false
  }
}
