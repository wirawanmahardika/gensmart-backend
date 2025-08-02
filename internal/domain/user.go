package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	GuruVolunteer *GuruVolunteer `json:"guru_volunteer,omitempty" gorm:"foreignKey:IDUser;references:ID"`
	Beasiswa      []Beasiswa     `json:"beasiswa,omitempty" gorm:"many2many:testimoni;foreignKey:ID;joinForeignKey:IDuser;references:ID;joinReferences:IDBeasiswa;"`
	Testimoni     *Testimoni     `json:"testimoni" gorm:"foreignKey:IDUser;references:ID;"`
}

func (e *Users) TableName() string { return "users" }
func (e *Users) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	return
}
