package products

import (
	"errors"
	"math"
	"net/http"

	"github.com/tolikcode/warehouse-api/db"
	"github.com/tolikcode/warehouse-api/features/articles"
	"github.com/tolikcode/warehouse-api/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type ProductArticleDto struct {
	ArticleID        int `json:"articleId"`
	QuantityRequired int `json:"quantityRequired"`
}

type ProductUpdateDto struct {
	Name            string              `json:"name"`
	ProductArticles []ProductArticleDto `json:"productArticles"`
}

type ProductDto struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Quantity int       `json:"quantity"`
}

// @Summary      Returns all products
// @Tags         products
// @Router       /products [get]
func GetProducts(c *gin.Context) {
	var products []Product
	err := db.DB.Model(&Product{}).Preload("ProductArticles.Article").Find(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productDtos := utils.MapSlice(products, func(p Product) ProductDto {
		productLimit := math.MaxInt

		if len(p.ProductArticles) == 0 {
			productLimit = 0
		}

		for _, pa := range p.ProductArticles {
			if pa.QuantityRequired == 0 {
				continue
			}

			articleLimit := pa.Article.Stock / pa.QuantityRequired
			if articleLimit < productLimit {
				productLimit = articleLimit
			}
		}

		return ProductDto{ID: p.ID, Name: p.Name, Quantity: productLimit}
	})

	c.JSON(http.StatusOK, gin.H{"data": productDtos})
}

// @Summary      Updates product definitions
// @Tags         products
// @Param products body string true "products"
// @Router       /products [patch]
func UpdateProducts(c *gin.Context) {
	var productUpdates []ProductUpdateDto
	if err := c.ShouldBindJSON(&productUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productNames := utils.MapSlice(productUpdates, func(a ProductUpdateDto) string { return a.Name })
	var existingProducts []Product
	db.DB.Where("name IN ?", productNames).Find(&existingProducts)

	updatedProducts := utils.MapSlice(productUpdates, func(pu ProductUpdateDto) Product {
		var product Product
		var existingIndex = slices.IndexFunc(existingProducts, func(p Product) bool { return p.Name == pu.Name })
		if existingIndex == -1 {
			product = Product{ID: uuid.New(), Name: pu.Name}
		} else {
			product = existingProducts[existingIndex]
		}
		product.ProductArticles = utils.MapSlice(pu.ProductArticles, func(pai ProductArticleDto) ProductArticle {
			return ProductArticle{ProductID: product.ID, ArticleID: pai.ArticleID, QuantityRequired: pai.QuantityRequired}
		})
		return product
	})

	updateResult := db.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(&updatedProducts)

	if updateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateResult.Error.Error()})
		return
	}

	deleteResult := db.DB.Not(map[string]interface{}{"name": productNames}).Delete(&Product{})
	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": deleteResult.Error.Error()})
		return
	}
}

// @Summary      Sells one specified product
// @Tags         products
// @Param        id   path      string  true  "Product ID"
// @Router       /products/{id}/sale [post]
func SellProduct(c *gin.Context) {
	var product Product
	if err := db.DB.Where("ID = ?", c.Param("id")).Preload("ProductArticles.Article").First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found!"})
		return
	}

	updateError := db.DB.Transaction(func(tx *gorm.DB) error {
		for _, pa := range product.ProductArticles {
			tx.Model(&articles.Article{}).Where("ID = ?", pa.ArticleID).Update("Stock", gorm.Expr("Stock - ?", pa.QuantityRequired))
			var article articles.Article
			tx.Model(&articles.Article{}).Where("ID = ?", pa.ArticleID).First(&article)
			if article.Stock < 0 {
				return errors.New("not enough in stock")
			}
		}

		return nil
	})

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateError.Error()})
		return
	}

}
