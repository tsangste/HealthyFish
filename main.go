package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tsangste/HealthyFish/controllers"
	"github.com/tsangste/HealthyFish/docs"
	"net/http"
)

// @title HealthyFish API
// @version 1.0
// @description HealthyFish item server.
// @contact.name Steven Tsang
// @contact.url http://github.com/tsangste/HealthyFish
// @contact.email tsangste@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/swagger/*any", func(context *gin.Context) {
		docs.SwaggerInfo.Host = context.Request.Host
		ginSwagger.CustomWrapHandler(&ginSwagger.Config{URL: "/swagger/doc.json"}, swaggerFiles.Handler)(context)
	})

	r.GET("/items", controllers.GetAll)
	r.GET("/items/calculate/:amount", controllers.CalculateItems)
	r.POST("/items/:item", controllers.AddItem)
	r.DELETE("/items/:item", controllers.DeleteItem)

	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
