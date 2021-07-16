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

	r.GET("/items/calculate/:id", controllers.CalculateItems)

	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
