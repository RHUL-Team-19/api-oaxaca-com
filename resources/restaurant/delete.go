package restaurant

import (
  "strconv"
  "github.com/gin-gonic/gin"
)

// deleteHandler is called when a HTTP DELETE request is sent to /restaurants.
func (r *Resource) deleteHandler(c *gin.Context) {
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

  // ...

}
