package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sekolah struct {
	ID               string    `json:"id"`
	IDUser           string    `json:"id_user"`
	Nama             string    `json:"nama"`
	Alamat           string    `json:"alamat"`
	StatusVerifikasi bool      `json:"status_verifikasi"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	User *Users `json:"user,omitempty" gorm:"foreignKey:IDUser;references:ID"`
}

func (s *Sekolah) TableName() string { return "sekolah" }
func (s *Sekolah) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}
