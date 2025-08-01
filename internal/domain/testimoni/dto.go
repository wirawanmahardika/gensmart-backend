package testimoniDomain

type CreateTestimoniRequest struct {
	IDUser     string
	IDBeasiswa string `json:"id_beasiswa"`
	Isi        string `json:"isi"`
}
