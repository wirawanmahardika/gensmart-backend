package usecase

import (
	"errors"
	"gensmart/internal/delivery/dto"
	"gensmart/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BeasiswaUsecase interface {
	Create(req *dto.CreateBeasiswaRequest) (err error)
	GetOne(id string) (beasiswa *domain.Beasiswa, err error)
	GetMany() (beasiswa []domain.Beasiswa, err error)
}

func NewBeasiswaUsecase(db *gorm.DB, validate *validator.Validate) BeasiswaUsecase {
	return &beasiswaUsecaseImpl{db, validate}
}

type beasiswaUsecaseImpl struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (uc *beasiswaUsecaseImpl) Create(req *dto.CreateBeasiswaRequest) (err error) {
	if err = uc.validate.Struct(req); err != nil {
		return
	}

	var countBeasiswa int64
	if err = uc.db.Model(&domain.Beasiswa{}).Where("nama = ?", req.Nama).Count(&countBeasiswa).Error; err != nil {
		return
	} else if countBeasiswa > 0 {
		return fiber.NewError(409, "beasiswa dengan nama tersebut sudah ada")
	}

	return uc.db.Create(&domain.Beasiswa{Nama: req.Nama, Link: req.Link, Deskripsi: req.Deskripsi}).Error
}

func (uc *beasiswaUsecaseImpl) GetOne(id string) (beasiswa *domain.Beasiswa, err error) {
	if err = uc.db.Where("id = ?", id).Take(&beasiswa).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fiber.NewError(404, "Beasiswa tidak ditemukan")
		}
		return
	}

	return
}

func (uc *beasiswaUsecaseImpl) GetMany() (beasiswa []domain.Beasiswa, err error) {
	if err = uc.db.Find(&beasiswa).Error; err != nil {
		return
	} else if len(beasiswa) == 0 {
		err = fiber.NewError(404, "beasiswa tidak ditemukan")
		return
	}

	return
}
