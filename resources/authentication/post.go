package authentication

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// postHandler is called when a HTTP POST request is made to /authentication.
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

  // fetch user by ID
  // ...

  // make sure passwords hashes match
  // ...

  // generate JWT
  // ...

}
