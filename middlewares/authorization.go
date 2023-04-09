package middlewares

import (
	"example_middleware/database"
	"example_middleware/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "invalid paramter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}
		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err !=nil{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":"Data Not Found",
				"message":"data doest exist",
			})
			return
		}
		if Product.UserID != userID{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":"unauthorized",
				"message":"you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
