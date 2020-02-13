package server

import (
  "github.com/gin-gonic/gin"
)

// corsMW is a middleware func that appends access-control headers to every
// incoming request.
func corsMW(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  c.Next()
}
