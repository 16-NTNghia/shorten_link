package services

import (
	"demo/internal/interfaces"
	"demo/internal/models"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

type LinkService struct {
	repo interfaces.LinkRepository
}

func NewLinkService(r interfaces.LinkRepository) *LinkService {
	return &LinkService{
		repo: r,
	}
}

func (ls *LinkService) GetAll() ([]*models.Link, error) {
	return ls.repo.FindAll()
}

func (ls *LinkService) GetByID(id uuid.UUID) (*models.Link, error) {
	return ls.repo.FindByID(id)
}

func (ls *LinkService) GetByCode(code string) (*models.Link, error) {
	return ls.repo.FindByCode(code)
}

func (ls *LinkService) CreateNewLink(url string) (*models.Link, error) {

	if !IsValidURL(url) {
		return nil, fmt.Errorf("invalid url")
	}

	l := models.Link{
		Url: url,
	}

	return ls.repo.Save(&l)
}

// IsValidURL kiểm tra xem một chuỗi có phải là URL hợp lệ hay không
func IsValidURL(s string) bool {
	parsedURL, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}

	// Kiểm tra xem có scheme (http/https) và host không
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}
	if parsedURL.Host == "" {
		return false
	}

	// Kiểm tra có ít nhất 1 dấu chấm trong host (tên miền hợp lệ)
	if !strings.Contains(parsedURL.Host, ".") && !strings.Contains(parsedURL.Host, "localhost") {
		return false
	}

	return true
}
