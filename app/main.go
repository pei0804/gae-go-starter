package server

import (
	"flag"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tikasan/gae-go-starter/api"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	echo *echo.Echo
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() {
	api := &api.Request{}
	s.echo = echo.New()

	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.CORS())

	s.echo.GET("/api/get", api.GetTest)
	s.echo.POST("/api/post", api.PostTest)
	s.echo.PUT("/api/put", api.PutTest)
	s.echo.DELETE("/api/delete", api.DeleteTest)

	s.echo.Pre(middleware.RemoveTrailingSlash())
	http.Handle("/", s.echo)
}

func init() {
	flag.Parse()
	s := New()
	s.Run()
}
