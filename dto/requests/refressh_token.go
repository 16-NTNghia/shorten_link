package requests

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
