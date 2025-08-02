package usecase

import (
	"errors"
	"gensmart/config"
	"gensmart/internal/delivery/dto"
	"gensmart/internal/domain"
	"gensmart/pkg"
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserUsecase interface {
	Register(req *dto.UserRegisterRequest) (err error)
	Login(req *dto.UserLoginRequest) (token string, err error)
	Data(email string) (user domain.Users, err error)
	GuruVolunteerUpdateStatusVerify(req *dto.GuruVolunteerUpdateStatusVerifyRequest) (err error)
}

func NewUserUsecase(db *gorm.DB, validate *validator.Validate) UserUsecase {
	return &userUsecaseImpl{db, validate}
}

type userUsecaseImpl struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (uc *userUsecaseImpl) Register(req *dto.UserRegisterRequest) (err error) {
	if err = uc.validate.Struct(req); err != nil {
		return
	}

	var countUser int64
	if err = uc.db.Model(&domain.Users{}).Where("email = ?", req.Email).Count(&countUser).Error; err != nil {
		return
	} else if countUser > 0 {
		return fiber.NewError(403, "email sudah terpakai, mohon gunakan email lain")
	}

	if !slices.Contains(config.Roles, req.Role) {
		return fiber.NewError(403, "pilih role yang sesuai")
	}

	hashedPassword, err := pkg.BcryptHash(req.Password)
	if err != nil {
		return fiber.NewError(403, "Terjadi kesalahan saat memproses password")
	}

	var user = &domain.Users{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Password: hashedPassword,
	}
	switch req.Role {
	case "guru_volunteer":
		user.GuruVolunteer = &domain.GuruVolunteer{
			Biodata:          req.Biodata,
			SertifikatUrl:    req.SertifikaURL,
			StatusVerifikasi: "pending",
		}
	case "admin_sekolah":
		user.Sekolah = &domain.Sekolah{
			Nama:             req.NamaSekolah,
			Alamat:           req.AlamatSekolah,
			StatusVerifikasi: false,
		}
	}

	if err = uc.db.Create(&user).Error; err != nil {
		return
	}
	return
}

func (uc *userUsecaseImpl) Login(req *dto.UserLoginRequest) (token string, err error) {
	if err = uc.validate.Struct(req); err != nil {
		return
	}

	var user domain.Users
	if err = uc.db.Where("email = ?", req.Email).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fiber.NewError(401, "email tidak terdaftar")
		}
		return
	} else if err = pkg.BcryptCompare(user.Password, req.Password); err != nil {
		err = fiber.NewError(401, "password salah")
		return
	}

	token, err = pkg.CreateJWTToken(user.ID, user.Role)
	if err != nil {
		return
	}
	return
}

func (uc *userUsecaseImpl) Data(id string) (user domain.Users, err error) {
	if err = uc.db.Preload("GuruVolunteer").Where("id = ?", id).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fiber.NewError(404, "data tidak ditemukan")
		}
		return
	}
	return
}

func (uc *userUsecaseImpl) GuruVolunteerUpdateStatusVerify(req *dto.GuruVolunteerUpdateStatusVerifyRequest) (err error) {
	var countUser int64
	if err = uc.db.Model(&domain.Users{}).Where("id = ? AND role = ?", req.IDUser, "guru_volunteer").Count(&countUser).Error; err != nil {
		return
	} else if countUser == 0 {
		return fiber.NewError(401, "guru volunteer yang dimaksud, tidak ditemukan")
	}

	return uc.db.Model(&domain.GuruVolunteer{}).Where("id_user = ?", req.IDUser).UpdateColumn("status_verifikasi", req.Status).Error
}
