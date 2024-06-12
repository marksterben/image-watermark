package server

import (
	"image-watermark/handler"
	"image-watermark/usecase"
	"log"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

const (
	APP_NAME = "mp-svc-catalog"
)

type Handler interface {
	AddWatermark(e echo.Context) error
}

type Server struct {
	e       *echo.Echo
	handler Handler
}

func NewServer() *Server {
	usecase := &usecase.Usecase{
		ContextTimeout: time.Duration(viper.GetInt("TIMEOUT")),
	}

	srv := &Server{
		e: echo.New(),
		handler: &handler.Handler{
			Usecase: usecase,
		},
	}

	return srv
}

func (s *Server) Run() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORS())
	s.e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "metrics") // Change "metrics" for your own path
		},
		Level: 5,
	}))

	s.routes()

	if err := s.e.Start(viper.GetString("PORT")); err != nil {
		log.Fatal(err)
	}
}
