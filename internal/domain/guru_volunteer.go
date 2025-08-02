package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GuruVolunteer struct {
	ID               string    `json:"id"`
	IDUser           string    `json:"id_user"`
	Biodata          string    `json:"biodata"`
	StatusVerifikasi string    `json:"status_verifikasi"`
	SertifikatUrl    string    `json:"sertifikat_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	User *Users `json:"user,omitempty" gorm:"foreignKey:IDUser;references:ID"`
}

func (s *GuruVolunteer) TableName() string { return "guru_volunteer" }
func (s *GuruVolunteer) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}
