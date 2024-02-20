package dto

type SignInRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponseDTO struct {
	SecretKey string `json:"secret_key"`
}
