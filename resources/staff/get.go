package staff

import (
  "strconv"
  "net/http"
  "github.com/gin-gonic/gin"
)

// getAllHandler is called when a HTTP GET request is made to /staff.
func (r *Resource) getAllHandler(c *gin.Context) {
  // fetch all from database
  staffs, err := r.db.GetAllStaff()
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // return result as JSON
  c.JSON(http.StatusOK, staffs)
}

// getOneHandler is called when a HTTP GET request is made to /staff/{id}.
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
  staff, err := r.db.GetStaff(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "ID does not exist"})
    return
  }

  // return result 200 as JSON
  c.JSON(http.StatusOK, staff)
}