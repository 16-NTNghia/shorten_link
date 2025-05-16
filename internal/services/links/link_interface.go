package links

import (
	"demo/internal/models"

	"github.com/google/uuid"
)

type Repository interface {
	FindAll() ([]*models.Link, error)
	FindByID(id uuid.UUID) (*models.Link, error)
	Save(e *models.Link) (*models.Link, error)
	FindByCode(code string) (*models.Link, error)
}
