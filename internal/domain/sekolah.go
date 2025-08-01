package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sekolah struct {
	ID               string    `json:"id"`
	Nama             string    `json:"nama"`
	Alamat           string    `json:"alamat"`
	StatusVerifikasi string    `json:"status_verifikasi"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (s *Sekolah) TableName() string { return "sekolah" }
func (s *Sekolah) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}
