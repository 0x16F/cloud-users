package httpsrv

import (
	"errors"
	"fmt"

	cerrors "github.com/0x16F/cloud-users/internal/usecase/errors"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	*fiber.App
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *cerrors.Error

	if errors.As(err, &e) {
		code = e.HttpCode
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return c.Status(code).SendString(err.Error())
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: errorHandler,
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
