package authentication

import (
  "github.com/dgrijalva/jwt-go"
)

// NewJWT generates a new JWT given some Claims.
func NewJWT(c jwt.Claims) *jwt.Token {
  return jwt.NewWithClaims(jwt.SigningMethodHS256, c)
}

// IsJWTStrValid takes a JWT string and a secret key string and tests whether
// the token is valid.
func IsJWTStrValid(t string, k string) (bool, error) {
  tkn, err := jwt.ParseWithClaims(
    t,
    &jwt.StandardClaims{},
    func(token *jwt.Token) (interface{}, error) {
      return []byte(k), nil
    },
  )
  if err != nil {
    return false, err
  } else {
    return tkn.Valid, nil
  }
}

// ParseJWTStr takes a JWT string and returns its StandardClaims. Does not
// perform validation and thus does not takes a secret key.
func ParseJWTStr(t string) (jwt.Claims, error) {
  if tkn, _, err := new(jwt.Parser).ParseUnverified(
    t,
    &jwt.StandardClaims{},
  ); err != nil {
    return nil, err
  } else {
    return tkn.Claims, nil
  }
}
