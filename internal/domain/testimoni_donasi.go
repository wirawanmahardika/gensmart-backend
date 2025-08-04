package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TestimoniDonasi struct {
	ID        string    `json:"id"`
	IDUser    string    `json:"id_users"`
	IDDonasi  string    `json:"id_donasi"`
	Isi       string    `json:"isi"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Users  *Users  `json:"user,omitempty" gorm:"foreignKey:IDUser;references:ID;"`
	Donasi *Donasi `json:"donasi,omitempty" gorm:"foreignKey:IDDonasi;references:ID;"`
}

func (e *TestimoniDonasi) TableName() string { return "testimoni_donasi" }
func (e *TestimoniDonasi) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	return
}
