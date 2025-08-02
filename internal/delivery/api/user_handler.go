package api

import (
	"gensmart/internal/delivery/dto"
	"gensmart/internal/usecase"
	"gensmart/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Register(c *fiber.Ctx) (err error)
	Login(c *fiber.Ctx) (err error)
	Data(c *fiber.Ctx) (err error)
	GuruVolunteerUpdateStatusVerify(c *fiber.Ctx) (err error)
	GetMany(c *fiber.Ctx) (err error)
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return &userHandlerImpl{uc}
}

type userHandlerImpl struct {
	uc usecase.UserUsecase
}

func (h *userHandlerImpl) Register(c *fiber.Ctx) (err error) {
	req := new(dto.UserRegisterRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	if err = h.uc.Register(req); err != nil {
		return
	}

	return c.Status(201).SendString("Berhasil daftar")
}

func (h *userHandlerImpl) Login(c *fiber.Ctx) (err error) {
	req := new(dto.UserLoginRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	token, err := h.uc.Login(req)
	if err != nil {
		return
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil login",
		"token":   token,
	})
}

func (h *userHandlerImpl) Data(c *fiber.Ctx) (err error) {
	id, _ := pkg.GetDataFromToken(c)
	user, err := h.uc.Data(id)
	if err != nil {
		return
	}

	return c.JSON(user)
}

func (h *userHandlerImpl) GuruVolunteerUpdateStatusVerify(c *fiber.Ctx) (err error) {
	req := new(dto.GuruVolunteerUpdateStatusVerifyRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	req.IDUser = c.Params("id")
	if err = h.uc.GuruVolunteerUpdateStatusVerify(req); err != nil {
		return
	}

	return c.SendString("Berhasil verifikasi guru volunteer")
}

func (h *userHandlerImpl) GetMany(c *fiber.Ctx) (err error) {
	users, err := h.uc.GetMany(c.Query("role"))
	if err != nil {
		return
	}

	return c.JSON(users)
}
