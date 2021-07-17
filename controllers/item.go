package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tsangste/HealthyFish/models"
	"github.com/tsangste/HealthyFish/services"
	http "net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	items := models.GetAll()

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func AddItem(c *gin.Context) {
	item := c.Param("item")

	i, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println(err)
	}

	models.Add(i)
	items := models.GetAll()

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func DeleteItem(c *gin.Context) {
	item := c.Param("item")

	i, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println(err)
	}

	models.Delete(i)
	items := models.GetAll()

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func CalculateItems(c *gin.Context) {
	amount := c.Param("amount")

	i, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println(err)
	}

	items := models.GetAll()
	orders := services.Calculate(items, i)

	c.JSON(http.StatusOK, gin.H{"data": orders})
}
