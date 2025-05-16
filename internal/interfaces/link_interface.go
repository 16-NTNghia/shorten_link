package interfaces

import (
	"demo/internal/models"

	"github.com/google/uuid"
)

type LinkRepository interface {
	FindAll() ([]*models.Link, error)
	FindByID(id uuid.UUID) (*models.Link, error)
	Save(e *models.Link) (*models.Link, error)
	FindByCode(code string) (*models.Link, error)
}

type LinkService interface {
	GetAll() ([]*models.Link, error)
	GetByID(id uuid.UUID) (*models.Link, error)
	GetByCode(code string) (*models.Link, error)
	CreateNewLink(url string) (*models.Link, error)
}