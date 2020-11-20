package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Base is parent controller
type Base struct{}

// response method
func (b *Base) response(c *fiber.Ctx, status int, message, data interface{}, code int) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"data":    data,
		"code":    code,
	})
}

// Success response success
func (b *Base) Success(c *fiber.Ctx, data interface{}) error {
	return b.response(c, http.StatusOK, "Success", data, http.StatusOK)
}

// Unauthorized response Unauthorized
func (b *Base) Unauthorized(c *fiber.Ctx) error {
	return b.response(c, http.StatusUnauthorized, "Unauthorized", nil, http.StatusUnauthorized)
}

// SuccessWithCode response success
func (b *Base) SuccessWithCode(c *fiber.Ctx, data interface{}, code int) error {
	return b.response(c, http.StatusOK, "Success", data, code)
}

// BadRequest response success
func (b *Base) BadRequest(c *fiber.Ctx, message string) error {
	return b.response(c, http.StatusBadRequest, message, nil, http.StatusBadRequest)
}
