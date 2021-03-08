package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wilgustavo/turnos/internal/domain/sucursal"
	"github.com/wilgustavo/turnos/internal/platform/server/handler"
)

// Server data
type Server struct {
	httpAddr string
	engine   *gin.Engine

	sucursalRepository sucursal.Repository
}

// New create a Server instance
func New(host string, port uint, sucursalRepository sucursal.Repository) Server {
	srv := Server{
		httpAddr:           fmt.Sprintf("%s:%d", host, port),
		engine:             gin.New(),
		sucursalRepository: sucursalRepository,
	}
	srv.registerRoutes()
	return srv
}

// Run runs the server
func (srv *Server) Run() error {
	log.Println("server running on", srv.httpAddr)
	return srv.engine.Run(srv.httpAddr)
}

func (srv *Server) registerRoutes() {
	srv.engine.POST("/sucursales", handler.CrearSucursalHandler(sucursal.NewCrearSucursalService(srv.sucursalRepository)))
}
