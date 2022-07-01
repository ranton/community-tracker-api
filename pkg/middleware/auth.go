package middleware

import (
	"strings"

	"github.com/VncntDzn/community-tracker-api/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const UNAUTHORIZED_MESSAGE string = "Unauthorized Request."

func AuthMiddleware(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": UNAUTHORIZED_MESSAGE})
	}

	authFields := strings.Fields(authHeader)
	if authFields[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": UNAUTHORIZED_MESSAGE})
	}

	token := authFields[1]
	_, tErr := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("JWT_SECRET")), nil
	})

	if tErr != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": UNAUTHORIZED_MESSAGE})
	}

	return ctx.Next()
}
