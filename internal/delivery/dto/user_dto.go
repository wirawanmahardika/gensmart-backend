package dto

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`

	// data untuk guru volunteer
	Biodata      string `json:"biodata"`
	SertifikaURL string `json:"sertifikat_url"`

	// data untuk admin sekolah
	NamaSekolah   string `json:"nama_sekolah"`
	AlamatSekolah string `json:"alamat_sekolah"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GuruVolunteerUpdateStatusVerifyRequest struct {
	IDUser string
	Status string `json:"status"`
}
