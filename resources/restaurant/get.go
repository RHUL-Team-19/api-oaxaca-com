package restaurant

import (
  "strconv"
  "net/http"
  "github.com/gin-gonic/gin"
)

// getAllHandler is called when a HTTP GET request is made to /restaurants.
func (r *Resource) getAllHandler(c *gin.Context) {
  // fetch all from database
  restaurants, err := r.db.GetAllRestaurants()
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // return result as JSON
  c.JSON(http.StatusOK, restaurants)
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

  // fetch one from database by ID
  restaurant, err := r.db.GetRestaurant(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "ID does not exist"})
    return
  }

  // return result 200 as JSON
  c.JSON(http.StatusOK, restaurant)
}
