package delivery

import (
	"gensmart/internal/delivery/api"
	"gensmart/internal/usecase"
	"gensmart/pkg"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(router *fiber.App, db *gorm.DB) {
	userUsecase := usecase.NewUserUsecase(db, pkg.Validate)
	userHandler := api.NewUserHandler(userUsecase)

	userRouter := router.Group("/v1/user")
	userRouter.Post("/register", userHandler.Register)
	userRouter.Post("/login", userHandler.Login)

	router.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	router.Get("/v1", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
}
