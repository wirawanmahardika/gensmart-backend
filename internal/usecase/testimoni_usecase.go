package usecase

import (
	"gensmart/internal/delivery/dto"
	"gensmart/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TestimoniUsecase interface {
	Create(req *dto.CreateTestimoniRequest) (err error)
	GetUsersTestimoniOnBeasiswa(id string) (users []domain.Users, err error)
}

func NewTestimoniUsecase(db *gorm.DB, validate *validator.Validate) TestimoniUsecase {
	return &testimoniUsecaseImpl{db, validate}
}

type testimoniUsecaseImpl struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (uc *testimoniUsecaseImpl) Create(req *dto.CreateTestimoniRequest) (err error) {
	if err = uc.validate.Struct(req); err != nil {
		return
	}

	var countBeasiswa int64
	if err = uc.db.Model(&domain.Beasiswa{}).Where("id = ?", req.IDBeasiswa).Count(&countBeasiswa).Error; err != nil {
		return
	} else if countBeasiswa == 0 {
		return fiber.NewError(404, "Beasiswa tidak ditemukan")
	}

	return uc.db.Create(&domain.Testimoni{
		IDUser:         req.IDUser,
		IDBeasiswa:     req.IDBeasiswa,
		Isi:            req.Isi,
		StatusModerasi: "pending",
	}).Error
}

func (uc *testimoniUsecaseImpl) GetUsersTestimoniOnBeasiswa(id string) (users []domain.Users, err error) {
	if err = uc.db.
		Model(&domain.Users{}).
		Joins("JOIN testimoni ON testimoni.id_user = users.id").
		Preload("Testimoni").
		Where("testimoni.id_beasiswa = ?", id).
		Find(&users).Error; err != nil {
		return
	}
	return
}
