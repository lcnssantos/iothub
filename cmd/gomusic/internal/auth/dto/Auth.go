package dto

type AuthRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	Type         string `json:"type"`
	RefreshToken string `json:"refreshToken"`
}
