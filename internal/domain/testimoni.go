package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Testimoni struct {
	ID             string    `json:"id"`
	IDUser         string    `json:"id_users"`
	IDBeasiswa     string    `json:"id_beasiswa"`
	Isi            string    `json:"isi"`
	StatusModerasi string    `json:"status_moderasi"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Users    *Users    `json:"user,omitempty" gorm:"foreignKey:IDUser;references:ID;"`
	Beasiswa *Beasiswa `json:"beasiswa,omitempty" gorm:"foreignKey:IDBeasiswa;references:ID;"`
}

func (e *Testimoni) TableName() string { return "testimoni" }
func (e *Testimoni) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	return
}
