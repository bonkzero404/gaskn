package middleware

import (
	"go-boilerplate-clean-arch/config"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Authenticate() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return utils.ApiUnauthorized(ctx, "Unauthorized user", err)
		},
		SigningKey: []byte(config.Config("JWT_SECRET")),
	})
}