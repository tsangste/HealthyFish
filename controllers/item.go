package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tsangste/HealthyFish/models"
	"github.com/tsangste/HealthyFish/services"
	http "net/http"
	"strconv"
)

// GetAll godoc
// @Summary Retrieves all item sizes
// @Produce json
// @success 200 {object} models.JSONResult{data=[]int}
// @Router /items [get]
func GetAll(c *gin.Context) {
	items := models.GetAll()

	c.JSON(http.StatusOK, gin.H{"data": items})
}

// CalculateItems godoc
// @Summary Get all the item sizes to fulfil size
// @Produce json
// @Param amount path int true "amount required"
// @success 200 {object} models.JSONResult{data=[]int}
// @Router /items/calculate/{amount} [get]
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

// AddItem godoc
// @Summary Adds item size
// @Produce json
// @Param item path int true "item size"
// @success 200 {object} models.JSONResult{data=[]int}
// @Router /items/{item} [post]
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

// DeleteItem godoc
// @Summary Deletes item size
// @Produce json
// @Param item path int true "item size"
// @success 200 {object} models.JSONResult{data=[]int}
// @Router /items/{item} [delete]
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
