package api

import (
	"gensmart/internal/delivery/dto"
	"gensmart/internal/usecase"
	"gensmart/pkg"

	"github.com/gofiber/fiber/v2"
)

type SekolahHandler interface {
	Create(c *fiber.Ctx) (err error)
	VerifikasiSekolah(c *fiber.Ctx) (err error)
	UpdateProfile(c *fiber.Ctx) (err error)
}

func NewSekolahHandler(uc usecase.SekolahUsecase) SekolahHandler {
	return &sekolahHandlerImpl{uc}
}

type sekolahHandlerImpl struct {
	uc usecase.SekolahUsecase
}

func (h *sekolahHandlerImpl) Create(c *fiber.Ctx) (err error) {
	req := new(dto.CreateSekolahRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	req.IDUser, _ = pkg.GetDataFromToken(c)
	if err = h.uc.Create(req); err != nil {
		return
	}

	return c.SendString("Berhasil membuat sekolah")
}

func (h *sekolahHandlerImpl) VerifikasiSekolah(c *fiber.Ctx) (err error) {
	req := new(dto.VerifikasiSekolahRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	req.IDSekolah = c.Params("id")
	if err = h.uc.VerifikasiSekolah(req); err != nil {
		return
	}

	return c.SendString("Berhasil verifikasi sekolah")
}

func (h *sekolahHandlerImpl) UpdateProfile(c *fiber.Ctx) (err error) {
	req := new(dto.UpdateProfileSekolahRequest)
	if err = c.BodyParser(req); err != nil {
		return pkg.BodyParserError()
	}

	req.IDUser, _ = pkg.GetDataFromToken(c)
	req.IDSekolah = c.Params("id")
	if err = h.uc.UpdateProfile(req); err != nil {
		return
	}

	return c.SendString("Berhasil update profile sekolah")
}
