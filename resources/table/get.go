package table

import (
  "strconv"
  "net/http"
  "github.com/gin-gonic/gin"
)

// getAllHandler is called when a HTTP GET request is made to /tables.
func (r *Resource) getAllHandler(c *gin.Context) {
  // get restaurant_id query param (empty string if not provided)
  restaurantID := c.Query("restaurant_id")

  // return all tables if no restaurant_id query param, else get all tables
  // for specific restaurant
  if restaurantID == "" {
    // fetch all from database
    tables, err := r.db.GetAllTables()
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    // return result as JSON
    c.JSON(http.StatusOK, tables)
  } else {
    // parse restaurant ID query param
    id, err := strconv.ParseInt(
      restaurantID,
      10,
      64,
    )
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Bad ID"})
      return
    }

    // fetch all from database
    tables, err := r.db.GetAllTablesForRestaurantID(id)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    // return result as JSON
    c.JSON(http.StatusOK, tables)
  }
}

// getAllForRestaurantHandler is called when a HTTP GET request is sent to
// /tables/restaurant/:id
func (r *Resource) getAllForRestaurantHandler(c *gin.Context) {
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

  // fetch all from database
  tables, err := r.db.GetAllTablesForRestaurantID(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // return result as JSON
  c.JSON(http.StatusOK, tables)
}

// getOneHandler is called when a HTTP GET request is made to /tables/{id}.
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
  table, err := r.db.GetTable(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "ID does not exist"})
    return
  }

  // return result 200 as JSON
  c.JSON(http.StatusOK, table)
}
