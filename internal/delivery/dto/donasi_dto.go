package dto

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
