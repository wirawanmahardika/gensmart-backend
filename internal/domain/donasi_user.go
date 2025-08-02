package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DonasiUser struct {
	ID        string    `json:"id"`
	IDDonasi  string    `json:"id_donasi"`
	IDUser    string    `json:"id_user"`
	Jumlah    int       `json:"jumlah"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Donasi *Donasi `json:"donasi,omitempty" gorm:"foreignKey:IDDonasi;references:ID"`
	User   *Users  `json:"user,omitempty" gorm:"foreignKey:IDUser;references:ID"`
}

func (e *DonasiUser) TableName() string { return "donasi_user" }
func (e *DonasiUser) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	return
}
