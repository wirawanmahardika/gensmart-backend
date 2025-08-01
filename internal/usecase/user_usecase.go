package usecase

import (
	"errors"
	"gensmart/config"
	"gensmart/internal/domain"
	userDomain "gensmart/internal/domain/user"
	"gensmart/pkg"
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserUsecase interface {
	Register(req *userDomain.UserRegisterRequest) (err error)
	Login(req *userDomain.UserLoginRequest) (token string, err error)
	Data(email string) (user domain.Users, err error)
}

func NewUserUsecase(db *gorm.DB, validate *validator.Validate) UserUsecase {
	return &userUsecaseImpl{db, validate}
}

type userUsecaseImpl struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (uc *userUsecaseImpl) Register(req *userDomain.UserRegisterRequest) (err error) {
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

	if err = uc.db.Create(&domain.Users{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Password: hashedPassword,
	}).Error; err != nil {
		return
	}

	return
}

func (uc *userUsecaseImpl) Login(req *userDomain.UserLoginRequest) (token string, err error) {
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

	token, err = pkg.CreateJWTToken(user.ID, user.Email)
	if err != nil {
		return
	}
	return
}

func (uc *userUsecaseImpl) Data(email string) (user domain.Users, err error) {
	if err = uc.db.Where("email = ?", email).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fiber.NewError(404, "data tidak ditemukan")
		}
		return
	}

	user.Password = ""
	return
}
