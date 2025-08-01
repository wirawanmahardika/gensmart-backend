package api

import (
	beasiswaDomain "gensmart/internal/domain/beasiswa"
	"gensmart/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type BeasiswaHandler interface {
	Create(c *fiber.Ctx) (err error)
	GetOne(c *fiber.Ctx) (err error)
}

func NewBeasiswaHandler(uc usecase.BeasiswaUsecase) BeasiswaHandler {
	return &BeasiswaHandlerImpl{uc}
}

type BeasiswaHandlerImpl struct {
	uc usecase.BeasiswaUsecase
}

func (h *BeasiswaHandlerImpl) Create(c *fiber.Ctx) (err error) {
	req := new(beasiswaDomain.CreateBeasiswaRequest)
	if err = c.BodyParser(req); err != nil {
		return
	}

	if err = h.uc.Create(req); err != nil {
		return
	}

	return c.SendString("Berhasil membuat beasiswa")
}

func (h *BeasiswaHandlerImpl) GetOne(c *fiber.Ctx) (err error) {
	beasiswa, err := h.uc.GetOne(c.Params("id"))
	if err != nil {
		return
	}

	return c.JSON(beasiswa)
}
