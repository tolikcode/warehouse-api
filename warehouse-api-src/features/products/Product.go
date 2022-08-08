package products

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID              uuid.UUID `json:"id" gorm:"type:uuid"`
	Name            string    `json:"name" gorm:"uniqueIndex"`
	ProductArticles []ProductArticle
}
