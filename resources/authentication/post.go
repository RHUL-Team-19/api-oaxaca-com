package authentication

import (
  "time"
  "strconv"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/dgrijalva/jwt-go"
  "api-oaxaca-com/packages/auth"
)

// postHandler is called when a HTTP POST request is made to /auth.
func (r *Resource) postHandler(c *gin.Context) {
  // unmarshaled JSON will be put here
  var input authenticationInput

  // parse input
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // input validation
  if valid, reason := input.IsValid(); !valid {
    c.JSON(http.StatusBadRequest, gin.H{"error": reason})
    return
  }

  // parse ID
  id, err := strconv.ParseInt(
    input.StaffID,
    10,
    64,
  )
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Bad ID"})
    return
  }

  // fetch staff by ID
  staff, err := r.db.GetStaff(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "ID does not exist"})
    return
  }

  // make sure password hashes match
  if m := auth.IsPasswordCorrect(
    input.Password,
    staff.PasswordHash,
  ); !m {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
    return
  }

  // generate JWT
  tkn := auth.NewJWT(
    &jwt.StandardClaims{
      Audience: "role goes here",
      IssuedAt: time.Now().Unix(),
      Subject: strconv.FormatInt(staff.StaffID, 10),
    },
  )
  if tknStr, err := tkn.SignedString([]byte("secret")); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
  } else {
    // return token string
    c.JSON(http.StatusOK, gin.H{"token": tknStr})
  }
}
