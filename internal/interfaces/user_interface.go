package interfaces

import (
	"demo/dto/requests"
	"demo/dto/responses"
	"demo/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll() ([]*models.User, error)
	FindByID(id uuid.UUID) (*models.User, error)
	Exists(id uuid.UUID) (bool, error)
	Save(e *models.User) (*models.User, error)
}

type UserService interface {
	GetAll() ([]*responses.UserResponse, error)
	GetByID(id uuid.UUID) (*responses.UserResponse, error)
	CreateNewUser(createUser *requests.CreateUserRequest) (*responses.UserResponse, error)
	UpdateUser(id uuid.UUID, updateUser *requests.UpdateUserRequest) (*responses.UserResponse, error)
}

type UserMapper interface {
	ToUserResponse(user *models.User) *responses.UserResponse
}
