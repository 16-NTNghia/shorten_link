package services

import (
	"demo/dto/requests"
	"demo/dto/responses"
	"demo/internal/configs"
	"demo/internal/interfaces"
	"demo/internal/models"
	"demo/pkg/middlewares"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo interfaces.UserRepository
}

func NewAuthService(repo interfaces.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

type MyClaims struct {
	jwt.RegisteredClaims
}

var secretKey string

func init() {
	configs.LoadEnv()

	secretKey = configs.GetEnv("JWT_SIGNER")
}

func CreateToken(userID string) (string, error) {
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "shorten_link",
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthService) Login(username string, password string) (*responses.AcceptTokenResponse, error) {

	if username == "" {
		return nil, errors.New("username is empty")
	}

	if password == "" {
		return nil, errors.New("password is empty")
	}

	fmt.Println(username, password)

	user, err := a.repo.FindByUsername(username)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	token, err := CreateToken(user.ID.String())

	if err != nil {
		return nil, err
	}

	return &responses.AcceptTokenResponse{AcceptToken: token}, nil
}

func (a *AuthService) RefreshToken(token string) (*responses.AcceptTokenResponse, error) {
	verifyToken, err := middlewares.VerifyToken(token)

	if err != nil {
		return nil, err
	}

	claims, ok := verifyToken.Claims.(jwt.MapClaims)

	if !ok || !verifyToken.Valid {
		return nil, errors.New("invalid token")
	}

	userID := claims["sub"].(string)

	token, err = CreateToken(userID)

	if err != nil {
		return nil, err
	}

	return &responses.AcceptTokenResponse{AcceptToken: token}, nil
}

func (a *AuthService) Register(register requests.RegisterRequest) (*responses.AcceptTokenResponse, error) {

	if register.Username == "" {
		return nil, errors.New("username is empty")
	}

	if register.Password == "" {
		return nil, errors.New("password is empty")
	}

	if register.Email == "" {
		return nil, errors.New("email is empty")
	}

	existedUser, err := a.repo.FindByUsername(register.Username)

	fmt.Println(existedUser)

	if err != nil {
		return nil, err
	}

	if existedUser != nil {
		return nil, errors.New("user already exists")
	}

	existedEmail, err := a.repo.ExistEmail(register.Email)

	if err != nil {
		return nil, err
	}

	if existedEmail {
		return nil, errors.New("email already exists")
	}

	registerUser := models.User{
		Username: register.Username,
		Password: register.Password,
		Email:    register.Email,
	}

	user, err := a.repo.Save(&registerUser)

	if err != nil {
		return nil, err
	}

	token, err := CreateToken(user.ID.String())

	if err != nil {
		return nil, err
	}

	return &responses.AcceptTokenResponse{AcceptToken: token}, nil
}
