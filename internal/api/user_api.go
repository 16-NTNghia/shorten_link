package api

import (
	"demo/dto/requests"
	"demo/dto/responses"
	"demo/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UsersHandler struct {
	service  interfaces.UserService
	validate *validator.Validate
}

func NewUsersHandler(s interfaces.UserService) *UsersHandler {
	return &UsersHandler{
		service:  s,
		validate: validator.New(),
	}
}

func (uh *UsersHandler) GetAll(c *gin.Context) {
	users, err := uh.service.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[[]*responses.UserResponse](err))
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse(users))
}

func (uh *UsersHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	user, err := uh.service.GetByID(uuid.MustParse(id))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*responses.UserResponse](err))
		return
	}

	c.IndentedJSON(http.StatusOK, responses.SuccessResponse(user))
}

func (uh *UsersHandler) CreateNewUser(c *gin.Context) {
	var newUser requests.CreateUserRequest
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*responses.UserResponse](err))
		return
	}

	if err := uh.validate.Struct(newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*responses.UserResponse](err))
		return
	}

	user, err := uh.service.CreateNewUser(&newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*responses.UserResponse](err))
		return
	}

	c.IndentedJSON(http.StatusOK, responses.SuccessResponse(user))
}

func (uh *UsersHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updateUser requests.UpdateUserRequest
	if err := c.BindJSON(&updateUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*responses.UserResponse](err))
		return
	}

	if err := uh.validate.Struct(updateUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*responses.UserResponse](err))
		return
	}

	user, err := uh.service.UpdateUser(uuid.MustParse(id), &updateUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*responses.UserResponse](err))
		return
	}

	c.IndentedJSON(http.StatusOK, responses.SuccessResponse(user))
}
