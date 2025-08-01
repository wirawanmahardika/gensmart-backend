package usecase

import (
	"errors"
	"gensmart/internal/delivery/dto"
	"gensmart/internal/domain"
	"gensmart/pkg"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DonasiUsecase interface {
	Create(req *dto.CreateDonasiRequest) (err error)
	UserDonate(req *dto.UserDonateRequest) (err error)
	VerifyUserDonate(req *dto.VerifyUserDonateRequest) (err error)
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
		Target:    req.Target,
		Progress:  0,
		Status:    "pending",
	}).Error
}

func (uc *donasiUsecaseImpl) UserDonate(req *dto.UserDonateRequest) (err error) {
	if err = uc.validate.Struct(req); err != nil {
		return
	}

	var countDonasi int64
	if err = uc.db.Model(&domain.Donasi{}).Where("id = ?", req.IDDonasi).Count(&countDonasi).Error; err != nil {
		return
	} else if countDonasi == 0 {
		return fiber.NewError(404, "tidak dapat melakukan donasi, tempat donasi yang dituju tidak tersedia")
	}

	return uc.db.Create(&domain.DonasiUser{
		IDDonasi: req.IDDonasi, IDUser: req.IDUser, Jumlah: req.Jumlah, Status: "pending",
	}).Error
}

func (uc *donasiUsecaseImpl) VerifyUserDonate(req *dto.VerifyUserDonateRequest) (err error) {
	if err = uc.validate.Struct(req); err != nil {
		return
	}

	return uc.db.Transaction(func(tx *gorm.DB) (err error) {
		var donasiUser domain.DonasiUser
		if err = tx.Take(&donasiUser, "id = ?", req.IDDonateUser).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fiber.NewError(404, "donasi dari user yang dimaksud tidak ditemukan")
			}
			return
		} else if donasiUser.Status != "pending" {
			return fiber.NewError(403, "donasi sudah terverikasi")
		} else if err = tx.Model(&domain.DonasiUser{}).Where("id = ?", req.IDDonateUser).UpdateColumn("status", req.Status).Error; err != nil {
			return
		} else if req.Status != "verified" {
			return
		}

		var donasi domain.Donasi
		if err = tx.Take(&donasi, "id = ?", donasiUser.IDDonasi).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fiber.NewError(404, "donasi tidak ditemukan")
			}
			return
		} else if donasi.Progress >= 100 {
			return fiber.NewError(403, "donasi sudah mencapai target")
		}

		donasi.Jumlah += donasiUser.Jumlah
		donasi.Progress = pkg.RoundToTwoDecimal(float64(donasi.Jumlah) / float64(donasi.Target))
		return tx.Select("jumlah", "progress").Updates(donasi).Error
	})
}
