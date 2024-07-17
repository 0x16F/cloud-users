package extractor

import (
	"github.com/0x16F/cloud-users/internal/entity"
	"github.com/gofiber/fiber/v2"
)

func Extract(c *fiber.Ctx) entity.UserData {
	login := string(c.Request().Header.Peek("CD_USER_LOGIN"))
	role := string(c.Request().Header.Peek("CD_USER_ROLE"))

	return entity.UserData{
		Login: login,
		Role:  role,
	}
}
