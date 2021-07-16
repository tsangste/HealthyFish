package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tsangste/HealthyFish/services"
	http "net/http"
	"strconv"
)

func CalculateItems(c *gin.Context) {
	item := c.Param("id")

	i, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println(err)
	}

	orders := services.Calculate(i)

	c.JSON(http.StatusOK, gin.H{"data": orders})
}
