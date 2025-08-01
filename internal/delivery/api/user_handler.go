package api

import (
	userDomain "gensmart/internal/domain/user"
	"gensmart/internal/usecase"
	"gensmart/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Register(c *fiber.Ctx) (err error)
	Login(c *fiber.Ctx) (err error)
	Data(c *fiber.Ctx) (err error)
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return &userHandlerImpl{uc}
}

type userHandlerImpl struct {
	uc usecase.UserUsecase
}

func (h *userHandlerImpl) Register(c *fiber.Ctx) (err error) {
	req := new(userDomain.UserRegisterRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	if err = h.uc.Register(req); err != nil {
		return
	}

	return c.Status(201).SendString("Berhasil daftar")
}

func (h *userHandlerImpl) Login(c *fiber.Ctx) (err error) {
	req := new(userDomain.UserLoginRequest)
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
	_, email := pkg.GetDataFromToken(c)
	user, err := h.uc.Data(email)
	if err != nil {
		return
	}

	return c.JSON(user)
}
