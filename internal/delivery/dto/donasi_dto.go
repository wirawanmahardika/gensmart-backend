package dto

type CreateDonasiRequest struct {
	IDUser    string
	IDSekolah string `json:"id_sekolah"`
	Jenis     string `json:"jenis"`
	Jumlah    int    `json:"jumlah"`
}
