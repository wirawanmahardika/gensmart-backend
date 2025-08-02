package middleware

import (
	"fmt"
	"gensmart/pkg"

	"github.com/gofiber/fiber/v2"
)

func RoleAuth(expectedRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		_, role := pkg.GetDataFromToken(c)
		fmt.Println(role)
		if expectedRole != role {
			return fiber.NewError(401, "anda tidak memiliki izin")
		}
		return c.Next()
	}
}
