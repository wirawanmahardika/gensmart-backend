package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Donasi struct {
	ID        string    `json:"id"`
	IDSekolah string    `json:"id_sekolah"`
	Jenis     string    `json:"jenis"`
	Jumlah    int       `json:"jumlah"`
	Target    int       `json:"target"`
	Status    string    `json:"status"`
	Progress  float64   `json:"progress"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User            *Users            `json:"user,omitempty" gorm:"foreignKey:IDSekolah;references:ID"`
	TestimoniDonasi []TestimoniDonasi `json:"testimoni_donasi,omitempty" gorm:"foreignKey:IDDonasi;references:id"`
	UserTestimoni   []Users           `json:"user_testimoni,omitempty" gorm:"many2many:testimoni_donasi;foreignKey:ID;joinForeignKey:IDDonasi;references:ID;joinReferences:IDUser"`
}

func (e *Donasi) TableName() string { return "donasi" }
func (e *Donasi) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	return
}
