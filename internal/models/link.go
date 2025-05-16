package models

import (
	"github.com/google/uuid"
	"time"
)

type Link struct {
	ID       uuid.UUID `json:"id"`
	Code     string    `json:"code"`
	Url      string    `json:"url"`
	CreateAt time.Time `json:"create_at"`
}
