package news

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	Id        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	AuthorId  uint           `json:"authorId"`
	View      int            `json:"view"`
}
