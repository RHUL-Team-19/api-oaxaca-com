package server

import (
  "github.com/gin-gonic/gin"
)

// type Server simply wraps a gin engine.
type Server struct {
  engine *gin.Engine
}

// Start takes a host interface and a port, and starts the server's gin engine
// listening on <host>:<port>.
func (s *Server) Start(host, port string) error {
  return s.engine.Run(host + ":" + port)
}
