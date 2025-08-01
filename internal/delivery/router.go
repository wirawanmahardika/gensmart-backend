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
	userRouter(router, db)
	beasiswaRouter(router, db)
}

func userRouter(router fiber.Router, db *gorm.DB) {
	userUsecase := usecase.NewUserUsecase(db, pkg.Validate)
	userHandler := api.NewUserHandler(userUsecase)

	userRouter := router.Group("/v1/user")
	userRouter.Post("/register", userHandler.Register)
	userRouter.Post("/login", userHandler.Login)

	userRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	userRouter.Get("/", userHandler.Data)
}

func beasiswaRouter(router fiber.Router, db *gorm.DB) {
	beasiswaUsecase := usecase.NewBeasiswaUsecase(db, pkg.Validate)
	beasiswaHandler := api.NewBeasiswaHandler(beasiswaUsecase)

	beasiswaRouter := router.Group("/v1/beasiswa")
	beasiswaRouter.Post("/", beasiswaHandler.Create)
	beasiswaRouter.Get("/:id", beasiswaHandler.GetOne)
	beasiswaRouter.Get("/", beasiswaHandler.GetMany)
}
