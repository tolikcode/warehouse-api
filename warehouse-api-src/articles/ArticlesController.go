package articles

import (
	"net/http"

	"github.com/tolikcode/warehouse-api/db"
	"github.com/tolikcode/warehouse-api/utils"

	"gorm.io/gorm/clause"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	var articles []Article
	db.DB.Find(&articles)

	c.JSON(http.StatusOK, gin.H{"data": articles})
}

func UpdateArticles(c *gin.Context) {
	var articlesUpdates []Article
	if err := c.ShouldBindJSON(&articlesUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateResult := db.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "stock"}),
	}).Create(&articlesUpdates)

	if updateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateResult.Error.Error()})
		return
	}

	articleIds := utils.MapSlice(articlesUpdates, func(a Article) int { return int(a.ID) })
	deleteResult := db.DB.Model(&Article{}).Not(map[string]interface{}{"id": articleIds}).Update("stock", 0)

	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": deleteResult.Error.Error()})
		return
	}
}
