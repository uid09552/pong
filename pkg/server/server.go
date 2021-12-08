package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"headerexplorer/pkg/headerExplorer"
)

const (
	DEFAULT_PORT = 5000
)

type Server struct{}

func (s *Server) Serve() {
	r := gin.Default()
	headerExplorer.Register(r)
	r.Run(fmt.Sprintf(":%d", DEFAULT_PORT))
}
