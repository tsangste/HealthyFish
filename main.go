package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tsangste/HealthyFish/controllers"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
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
