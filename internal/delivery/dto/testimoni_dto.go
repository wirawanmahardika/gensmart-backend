package dto

type CreateTestimoniRequest struct {
	IDUser     string
	IDBeasiswa string `json:"id_beasiswa"`
	Isi        string `json:"isi"`
}

type UpdateStatusTestimoniRequest struct {
	IDTestimoni string
	Status      string `json:"status"`
}
