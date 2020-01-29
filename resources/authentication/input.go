package authentication

import (
  "strconv"
)

// type authenticationInput contains fields send via a HTTP POST request to
// /authentication when requesting a new auth token.
type authenticationInput struct {
  StaffID string `json:"staff_id"`
  Password string `json:"password"`
}

// IsValid ensures the fields provided in authenticationInput match as set of
// rules. If validation fails, a reason is returned - which should be send to
// the user.
func (i *authenticationInput) IsValid() (bool, string) {
  if _, err := strconv.ParseInt(i.StaffID, 10, 64); err != nil {
    return false, "User ID must be numerical"
  }

  if len(i.Password) < 5 || len(i.Password) >= 64 {
    return false, "Length of password must be between 5 and 64"
  }

  return true, ""
}
