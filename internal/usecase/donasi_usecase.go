package usecase

import (
	"gensmart/internal/delivery/dto"
	"gensmart/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DonasiUsecase interface {
	Create(req *dto.CreateDonasiRequest) (err error)
}

func NewDonasiUsecase(db *gorm.DB, validate *validator.Validate) DonasiUsecase {
	return &donasiUsecaseImpl{db, validate}
}

type donasiUsecaseImpl struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (uc *donasiUsecaseImpl) Create(req *dto.CreateDonasiRequest) (err error) {
	if err = uc.validate.Struct(req); err != nil {
		return
	}

	var countUser int64
	if err = uc.db.Model(&domain.Sekolah{}).Where("id = ? AND id_user = ?", req.IDSekolah, req.IDUser).Count(&countUser).Error; err != nil {
		return
	} else if countUser == 0 {
		return fiber.NewError(401, "user tidak memiliki izin")
	}

	return uc.db.Create(&domain.Donasi{
		IDSekolah: req.IDSekolah,
		Jenis:     req.Jenis,
		Jumlah:    req.Jumlah,
		Progress:  0,
		Status:    "pending",
	}).Error
}
