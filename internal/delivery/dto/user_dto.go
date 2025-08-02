package dto

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`

	Biodata      string `json:"biodata"`
	SertifikaURL string `json:"sertifikat_url"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
