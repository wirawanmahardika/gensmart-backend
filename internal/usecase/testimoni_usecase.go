package usecase

import (
	beasiswaDomain "gensmart/internal/domain/beasiswa"
	testimoniDomain "gensmart/internal/domain/testimoni"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TestimoniUsecase interface {
	Create(req *testimoniDomain.CreateTestimoniRequest) (err error)
}

func NewTestimoniUsecase(db *gorm.DB, validate *validator.Validate) TestimoniUsecase {
	return &testimoniUsecaseImpl{db, validate}
}

type testimoniUsecaseImpl struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (uc *testimoniUsecaseImpl) Create(req *testimoniDomain.CreateTestimoniRequest) (err error) {
	if err = uc.validate.Struct(req); err != nil {
		return
	}

	var countBeasiswa int64
	if err = uc.db.Model(&beasiswaDomain.Entity{}).Where("id = ?", req.IDBeasiswa).Count(&countBeasiswa).Error; err != nil {
		return
	} else if countBeasiswa == 0 {
		return fiber.NewError(404, "Beasiswa tidak ditemukan")
	}

	return uc.db.Create(&testimoniDomain.Entity{
		IDUser:         req.IDUser,
		IDBeasiswa:     req.IDBeasiswa,
		Isi:            req.Isi,
		StatusModerasi: "pending",
	}).Error
}
