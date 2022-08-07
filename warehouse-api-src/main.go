package main

import (
	"net/http"

	"github.com/tolikcode/warehouse-api/db"
	"github.com/tolikcode/warehouse-api/docs"
	"github.com/tolikcode/warehouse-api/features/articles"
	"github.com/tolikcode/warehouse-api/features/products"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Warehouse API
func main() {
	docs.SwaggerInfo.Host = "localhost:8080"

	db.ConnectDatabase()
	db.DB.AutoMigrate(&articles.Article{}, &products.Product{}, &products.ProductArticle{})

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "warehouse-api is working") })
	r.GET("/articles", articles.GetArticles)
	r.PATCH("/articles", articles.UpdateArticles)
	r.GET("/products", products.GetProducts)
	r.PATCH("/products", products.UpdateProducts)
	r.POST("/products/:id/sale", products.SellProduct)

	r.Run()
}
