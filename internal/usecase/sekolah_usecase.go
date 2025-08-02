package usecase

import (
	"errors"
	"gensmart/internal/delivery/dto"
	"gensmart/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SekolahUsecase interface {
	Create(req *dto.CreateSekolahRequest) (err error)
	VerifikasiSekolah(req *dto.VerifikasiSekolahRequest) (err error)
}

func NewSekolahUsecase(db *gorm.DB, validate *validator.Validate) SekolahUsecase {
	return &sekolahUsecaseImpl{db, validate}
}

type sekolahUsecaseImpl struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (uc *sekolahUsecaseImpl) Create(req *dto.CreateSekolahRequest) (err error) {
	var countUser int64
	if err = uc.db.Model(&domain.Users{}).Where("id = ? AND role = ?", req.IDUser, "admin_sekolah").Count(&countUser).Error; err != nil {
		return
	} else if countUser == 0 {
		return fiber.NewError(401, "user tidak memiliki izin untuk membuat sekolah")
	}

	var countSekolah int64
	if err = uc.db.Model(&domain.Sekolah{}).Where("id_user = ?", req.IDUser).Count(&countSekolah).Error; err != nil {
		return
	} else if countSekolah > 0 {
		return fiber.NewError(401, "user sudah menjadi admin dari suatu sekolah")
	}

	return uc.db.Create(&domain.Sekolah{
		IDUser:           req.IDUser,
		Nama:             req.Nama,
		Alamat:           req.Alamat,
		StatusVerifikasi: false,
	}).Error
}

func (uc *sekolahUsecaseImpl) VerifikasiSekolah(req *dto.VerifikasiSekolahRequest) (err error) {
	var sekolah *domain.Sekolah
	if err = uc.db.Where("id = ?", req.IDSekolah).Take(&sekolah).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fiber.NewError(403, "sekolah yang ingin diverifikasi tidak ditemukan")
		}
		return
	}

	return uc.db.Model(&domain.Sekolah{}).Where("id = ?", req.IDSekolah).UpdateColumn("status_verifikasi", req.Status).Error
}
