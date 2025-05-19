package api

import (
	"demo/dto/requests"
	"demo/dto/responses"
	"demo/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service interfaces.AuthService
}

func NewAuthHandler(s interfaces.AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var Login requests.LoginRequest

	if err := c.BindJSON(&Login); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[string](err))
		return
	}

	token, err := ah.service.Login(Login.Username, Login.Password)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[string](err))
		return
	}

	c.IndentedJSON(http.StatusOK, responses.SuccessResponse(token))
}

func (ah *AuthHandler) RefreshToken(c *gin.Context) {
	var token requests.RefreshTokenRequest

	if err := c.BindJSON(&token); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[string](err))
		return
	}

	result, err := ah.service.RefreshToken(token.RefreshToken)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[string](err))
		return
	}

	c.IndentedJSON(http.StatusOK, responses.SuccessResponse(result))
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var register requests.RegisterRequest

	if err := c.BindJSON(&register); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[string](err))
		return
	}

	result, err := ah.service.Register(register)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[string](err))
		return
	}

	c.IndentedJSON(http.StatusOK, responses.SuccessResponse(result))
}
