package httpsrv

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	*fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	return &Server{
		App: app,
	}
}

func (s *Server) Start(port uint16) error {
	s.App.Use(logger.New())

	return s.App.Listen(fmt.Sprintf(":%d", port))
}

func (s *Server) Stop() error {
	return s.App.Shutdown()
}
