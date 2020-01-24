package restaurant

import (
  "strconv"
  "net/http"
  "github.com/gin-gonic/gin"
)

// putHandler is called when a HTTP PUT request is made to /restaurants.
func (r *Resource) putHandler(c *gin.Context) {
  // parse ID from URL
  id, err := strconv.ParseInt(
    c.Param("id"),
    10,
    64,
  )
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Bad ID"})
    return
  }

  // unmarshaled JSON will be put here
  var input updateInput

  // parse input
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // input validation
  // ...

  // pass input to DB method
  if err := r.db.UpdateRestaurant(id, &input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // return success
  c.Status(http.StatusCreated)
}
