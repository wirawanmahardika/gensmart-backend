package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Beasiswa struct {
	ID        string    `json:"id"`
	Nama      string    `json:"nama"`
	Link      string    `json:"link"`
	Deskripsi string    `json:"deskripsi"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Users []Users `json:"users,omitempty" gorm:"many2many:testimoni;foreignKey:ID;joinForeignKey:IDBeasiswa;references:ID;joinReferences:IDUser;"`
}

func (e *Beasiswa) TableName() string { return "beasiswa" }
func (e *Beasiswa) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	return
}
