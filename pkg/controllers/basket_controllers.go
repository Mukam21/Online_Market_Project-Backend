package controllers

import (
	"net/http"
	"online_market_project/pkg/database"
	"online_market_project/pkg/models"

	"github.com/gin-gonic/gin"
)

func AddToProduc(c *gin.Context) {
	var cart models.Produc
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, cart.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := database.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
		return
	}

	c.JSON(http.StatusCreated, cart)
}

func GetProduc(c *gin.Context) {
	userID := c.Param("user_id")
	var cartItems []models.Produc

	if err := database.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

func RemoveFromProduc(c *gin.Context) {
	cartID := c.Param("id")

	if err := database.DB.Delete(&models.Produc{}, cartID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
