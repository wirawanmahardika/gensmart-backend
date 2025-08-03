package api

import (
	"gensmart/internal/delivery/dto"
	"gensmart/internal/usecase"
	"gensmart/pkg"

	"github.com/gofiber/fiber/v2"
)

type DonasiHandler interface {
	Create(c *fiber.Ctx) (err error)
	UserDonate(c *fiber.Ctx) (err error)
	VerifyUserDonate(c *fiber.Ctx) (err error)
	GetOne(c *fiber.Ctx) (err error)
}

func NewDonasiHandler(uc usecase.DonasiUsecase) DonasiHandler {
	return &donasiHandlerImpl{uc}
}

type donasiHandlerImpl struct {
	uc usecase.DonasiUsecase
}

func (h *donasiHandlerImpl) Create(c *fiber.Ctx) (err error) {
	req := new(dto.CreateDonasiRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	req.IDUser, _ = pkg.GetDataFromToken(c)
	if err = h.uc.Create(req); err != nil {
		return
	}

	return c.SendString("Berhasil mengajukan permintaan donasi")
}

func (h *donasiHandlerImpl) UserDonate(c *fiber.Ctx) (err error) {
	req := new(dto.UserDonateRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	req.IDUser, _ = pkg.GetDataFromToken(c)
	if err = h.uc.UserDonate(req); err != nil {
		return
	}

	return c.SendString("Berhasil melakukan donasi, silahkan tunggu verifikasi")
}

func (h *donasiHandlerImpl) VerifyUserDonate(c *fiber.Ctx) (err error) {
	req := new(dto.VerifyUserDonateRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	req.IDDonateUser = c.Params("id")
	if err = h.uc.VerifyUserDonate(req); err != nil {
		return
	}

	return c.SendString("Berhasil verifikasi donasi")
}

func (h *donasiHandlerImpl) GetOne(c *fiber.Ctx) (err error) {
	donasi, err := h.uc.GetOne(c.Params("id"))
	if err != nil {
		return
	}

	return c.JSON(donasi)
}
