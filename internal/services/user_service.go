package services

import (
	"demo/dto/requests"
	"demo/dto/responses"
	"demo/internal/interfaces"
	"demo/internal/models"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo   interfaces.UserRepository
	mapper interfaces.UserMapper
}

func NewUserService(r interfaces.UserRepository, m interfaces.UserMapper) *UserService {
	return &UserService{
		repo:   r,
		mapper: m,
	}
}

func (u *UserService) GetAll() ([]*responses.UserResponse, error) {
	listUser, err := u.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var response []*responses.UserResponse

	for _, user := range listUser {
		response = append(response, u.mapper.ToUserResponse(user))
	}

	return response, nil
}

func (u *UserService) GetByID(id uuid.UUID) (*responses.UserResponse, error) {
	user, err := u.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	response := u.mapper.ToUserResponse(user)

	return response, nil
}

func (u *UserService) CreateNewUser(createUser *requests.CreateUserRequest) (*responses.UserResponse, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Username: createUser.Username,
		Password: string(hashPassword),
		Email:    createUser.Email,
	}

	user, err := u.repo.Save(&newUser)

	if err != nil {
		return nil, err
	}

	response := u.mapper.ToUserResponse(user)

	return response, nil
}

func (u *UserService) UpdateUser(id uuid.UUID, updateUser *requests.UpdateUserRequest) (*responses.UserResponse, error) {
	modifyUser, err := u.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	if strings.Contains(updateUser.Username, "") {
		modifyUser.Username = updateUser.Username
	}

	if strings.Contains(updateUser.Email, "") {
		modifyUser.Email = updateUser.Email
	}

	user, err := u.repo.Save(modifyUser)

	if err != nil {
		return nil, err
	}

	response := u.mapper.ToUserResponse(user)

	return response, nil
}
