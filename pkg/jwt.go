package pkg

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(id, email string) (token string, err error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaim.SignedString([]byte("secret"))
}

func GetDataFromToken(c *fiber.Ctx) (id string, email string) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id = claims["id"].(string)
	email = claims["email"].(string)
	return
}
