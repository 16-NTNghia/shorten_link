package interfaces

import (
	"demo/dto/requests"
	"demo/dto/responses"
)

type AuthService interface {
	Login(username string, password string) (*responses.AcceptTokenResponse, error)
	RefreshToken(token string) (*responses.AcceptTokenResponse, error)
	Register(register requests.RegisterRequest) (*responses.AcceptTokenResponse, error)
}
