package dto

type CreateSekolahRequest struct {
	IDUser string
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
}

type VerifikasiSekolahRequest struct {
	IDSekolah string
	Status    bool `json:"status"`
}

type UpdateProfileSekolahRequest struct {
	IDUser    string
	IDSekolah string
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
}
