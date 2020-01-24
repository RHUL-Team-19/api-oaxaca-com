package restaurant

import (
  "strconv"
  "github.com/gin-gonic/gin"
)

// getAllHandler is called when a HTTP GET request is made to /restaurants.
func (r *Resource) getAllHandler(c *gin.Context) {
  // ...
  restaurants, err := r.db.GetAllResturants()
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // ...

}

// getOneHandler is called when a HTTP GET request is made to /restaurants/{id}.
func (r *Resource) getOneHandler(c *gin.Context) {
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
  restaurant, err := r.db.GetResturant(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "ID does not exist"})
    return
  }

  // ...

}
