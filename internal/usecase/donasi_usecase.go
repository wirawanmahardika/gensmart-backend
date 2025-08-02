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
	UserDonate(req *dto.UserDonateRequest) (err error)
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

// func (uc *donasiUsecaseImpl) VerifyUserDonate(req *dto.UserDonateRequest) (err error) {
// 	if err = uc.validate.Struct(req); err != nil {
// 		return
// 	}

// 	var donasi domain.Donasi
// 	if err = uc.db.Model(&domain.Donasi{}).Where("id = ?", req.IDDonasi).Take(&donasi).Error; err != nil {
// 		return
// 	}

// 	return uc.db.Transaction(func(tx *gorm.DB) (err error) {
// 		donasi.Jumlah += req.Jumlah
// 		donasi.Progress = pkg.RoundToTwoDecimal(float64(donasi.Jumlah / donasi.Target))

// 		if err = tx.Create(&domain.DonasiUser{IDDonasi: req.IDDonasi, IDUser: req.IDUser, Jumlah: req.Jumlah, Status: "pending"}).Error; err != nil {
// 			return
// 		}
// 		return
// 	})
// }
