package domain

import (
	"encoding/json"
	"time"
)

type Article struct {
	Id        int64           `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	AuthorId  string          `json:"author_id"`
	CreatedAt time.Time       `gorm:"<-:create;" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	Tag       json.RawMessage `json:"tag,omitempty"`
}
