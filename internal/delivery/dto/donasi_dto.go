package dto

import "time"

type CreateDonasiRequest struct {
	IDUser    string
	IDSekolah string `json:"id_sekolah"`
	Jenis     string `json:"jenis"`
	Target    int    `json:"target"`
}

type VerifyDonateRequest struct {
	IDDonate string
	Status   string `json:"status"`
}

type UserDonateRequest struct {
	IDUser   string
	IDDonasi string `json:"id_donasi"`
	Jumlah   int    `json:"jumlah"`
}

type VerifyUserDonateRequest struct {
	IDDonateUser string
	Status       string `json:"status"`
}

type UserTestimoniDonationRequest struct {
	IDDonasi string
	IDUser   string
	Isi      string `json:"isi"`
}

// response

type GetOneDonasiResponse struct {
	ID        string    `json:"id"`
	IDSekolah string    `json:"id_sekolah"`
	Jenis     string    `json:"jenis"`
	Jumlah    int       `json:"jumlah"`
	Target    int       `json:"target"`
	Status    string    `json:"status"`
	Progress  float64   `json:"progress"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TestimoniDonasi []Donatur `json:"testimoni_donasi,omitempty"`
}
type Donatur struct {
	Donatur      string `json:"donatur"`
	EmailDonatur string `json:"email_donatur"`
	RoleDonatur  string `json:"role_donatur"`
	Pesan        string `json:"pesan"`
}
