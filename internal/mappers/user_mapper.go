package mappers

import (
	"demo/dto/responses"
	"demo/internal/models"
)

type DefaultUserMapper struct{}

func NewUserMapper() *DefaultUserMapper {
	return &DefaultUserMapper{}
}

func (m *DefaultUserMapper) ToUserResponse(user *models.User) *responses.UserResponse {
	return &responses.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Actived:  user.Actived,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
}
