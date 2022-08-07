package products

import (
	"github.com/tolikcode/warehouse-api/articles"

	"github.com/google/uuid"
)

type ProductArticle struct {
	ProductID        uuid.UUID        `gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	ArticleID        int              `gorm:"primaryKey"`
	Article          articles.Article `gorm:"foreignKey:ArticleID"`
	QuantityRequired int
}
