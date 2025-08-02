package api

import (
	"gensmart/internal/delivery/dto"
	"gensmart/internal/usecase"
	"gensmart/pkg"

	"github.com/gofiber/fiber/v2"
)

type TestimoniHandler interface {
	Create(c *fiber.Ctx) (err error)
	GetUsersTestimoniOnBeasiswa(c *fiber.Ctx) (err error)
	UpdateStatusTestimoni(c *fiber.Ctx) (err error)
}

func NewTestimoniHandler(uc usecase.TestimoniUsecase) TestimoniHandler {
	return &testimoniHandlerImpl{uc}
}

type testimoniHandlerImpl struct {
	uc usecase.TestimoniUsecase
}

func (h *testimoniHandlerImpl) Create(c *fiber.Ctx) (err error) {
	req := new(dto.CreateTestimoniRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	id, _ := pkg.GetDataFromToken(c)
	req.IDUser = id
	if err = h.uc.Create(req); err != nil {
		return
	}

	return c.SendString("Berhasil memberi testimoni")
}

func (h *testimoniHandlerImpl) GetUsersTestimoniOnBeasiswa(c *fiber.Ctx) (err error) {
	users, err := h.uc.GetUsersTestimoniOnBeasiswa(c.Params("id"))
	if err != nil {
		return
	}

	return c.JSON(users)
}
func (h *testimoniHandlerImpl) UpdateStatusTestimoni(c *fiber.Ctx) (err error) {
	req := new(dto.UpdateStatusTestimoniRequest)
	if err = c.BodyParser(req); err != nil {
		return
	}

	req.IDTestimoni = c.Params("id")
	if err = h.uc.UpdateStatusTestimoni(req); err != nil {
		return
	}

	return c.SendString("Berhasil mengubah status testimoni")
}
