package api

import (
	testimoniDomain "gensmart/internal/domain/testimoni"
	"gensmart/internal/usecase"
	"gensmart/pkg"

	"github.com/gofiber/fiber/v2"
)

type TestimoniHandler interface {
	Create(c *fiber.Ctx) (err error)
}

func NewTestimoniHandler(uc usecase.TestimoniUsecase) TestimoniHandler {
	return &testimoniHandlerImpl{uc}
}

type testimoniHandlerImpl struct {
	uc usecase.TestimoniUsecase
}

func (h *testimoniHandlerImpl) Create(c *fiber.Ctx) (err error) {
	req := new(testimoniDomain.CreateTestimoniRequest)
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
