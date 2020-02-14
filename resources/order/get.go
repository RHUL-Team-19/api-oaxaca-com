package order

import (
  "strconv"
  "net/http"
  "github.com/gin-gonic/gin"
)

// getAllHandler is called when a HTTP GET request is made to /orders.
func (r *Resource) getAllHandler(c *gin.Context) {
  // get restaurant_id query param (empty string if not provided)
  restaurantID := c.Query("restaurant_id")

  // get table_id query param (empty string if not provided)
  tableID := c.Query("table_id")

  // return all orders if no restaurant_id query param and no table_id query
  // param, else get all orders for a specific table, else get all orders for
  // a specific restaurant, else return all orders for a specific table
  if tableID != "" {
    // parse table ID query param
    id, err := strconv.ParseInt(
      tableID,
      10,
      64,
    )
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Bad ID"})
      return
    }

    // fetch all from database
    orders, err := r.db.GetAllOrdersForTableID(id)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    // return result as JSON
    c.JSON(http.StatusOK, orders)
  } else if restaurantID != "" {
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
    orders, err := r.db.GetAllOrdersForRestaurantID(id)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    // return result as JSON
    c.JSON(http.StatusOK, orders)
  } else {
    // fetch all from database
    orders, err := r.db.GetAllOrders()
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    // return result as JSON
    c.JSON(http.StatusOK, orders)
  }
}

// getOneHandler is called when a HTTP GET request is made to /orders/{id}.
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
  order, err := r.db.GetOrder(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "ID does not exist"})
    return
  }

  // return result 200 as JSON
  c.JSON(http.StatusOK, order)
}

// getMealsForOneHandler is called when a HTTP GET request is made to
// /orders/{id}/meals.
func (r *Resource) getMealsForOneHandler(c *gin.Context) {
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

  // fetch all from database by ID
  meals, err := r.db.GetMealsForOrderID(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "ID does not exist"})
    return
  }

  // return result 200 as JSON
  c.JSON(http.StatusOK, meals)
}
