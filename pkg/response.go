package pkg

import "github.com/gofiber/fiber/v2"

type RegularResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func BodyParserError() error {
	return fiber.NewError(400, "Gagal melakukan parsing pada body request")
}
