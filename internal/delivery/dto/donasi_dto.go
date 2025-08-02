package dto

type CreateDonasiRequest struct {
	IDUser    string
	IDSekolah string `json:"id_sekolah"`
	Jenis     string `json:"jenis"`
	Target    int    `json:"target"`
}

type UserDonateRequest struct {
	IDUser   string
	IDDonasi string `json:"id_donasi"`
	Jumlah   int    `json:"jumlah"`
}
