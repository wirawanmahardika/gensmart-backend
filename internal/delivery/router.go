package delivery

import (
	"gensmart/internal/delivery/api"
	"gensmart/internal/delivery/middleware"
	"gensmart/internal/usecase"
	"gensmart/pkg"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(router *fiber.App, db *gorm.DB) {
	userRouter(router, db)
	beasiswaRouter(router, db)
	testimoniRouter(router, db)
	sekolahRouter(router, db)
	donasiRouter(router, db)
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
	userRouter.Patch("/guru-volunteer/:id/status", middleware.RoleAuth("admin"), userHandler.GuruVolunteerUpdateStatusVerify)
}

func beasiswaRouter(router fiber.Router, db *gorm.DB) {
	beasiswaUsecase := usecase.NewBeasiswaUsecase(db, pkg.Validate)
	beasiswaHandler := api.NewBeasiswaHandler(beasiswaUsecase)

	beasiswaRouter := router.Group("/v1/beasiswa")
	beasiswaRouter.Post("/", beasiswaHandler.Create)
	beasiswaRouter.Get("/:id", beasiswaHandler.GetOne)
	beasiswaRouter.Get("/", beasiswaHandler.GetMany)
}

func testimoniRouter(router fiber.Router, db *gorm.DB) {
	testimoniUsecase := usecase.NewTestimoniUsecase(db, pkg.Validate)
	testimoniHandler := api.NewTestimoniHandler(testimoniUsecase)

	testimoniRouter := router.Group("/v1/testimoni")
	testimoniRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	testimoniRouter.Post("/", testimoniHandler.Create)
	testimoniRouter.Get("/beasiswa/:id", testimoniHandler.GetUsersTestimoniOnBeasiswa)
	testimoniRouter.Patch("/beasiswa/:id/status", middleware.RoleAuth("admin"), testimoniHandler.UpdateStatusTestimoni)
}

func sekolahRouter(router fiber.Router, db *gorm.DB) {
	sekolahUsecase := usecase.NewSekolahUsecase(db, pkg.Validate)
	sekolahHandler := api.NewSekolahHandler(sekolahUsecase)

	sekolahRouter := router.Group("/v1/sekolah")
	sekolahRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	sekolahRouter.Post("/", sekolahHandler.Create)
	sekolahRouter.Patch("/:id", sekolahHandler.VerifikasiSekolah)
}

func donasiRouter(router fiber.Router, db *gorm.DB) {
	donasiUsecase := usecase.NewDonasiUsecase(db, pkg.Validate)
	donasiHandler := api.NewDonasiHandler(donasiUsecase)

	donasiRouter := router.Group("/v1/donasi")
	donasiRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	donasiRouter.Post("/", donasiHandler.Create)
	donasiRouter.Post("/user", donasiHandler.UserDonate)
}
