package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/bayunashr/gopoll/initializers"
	"github.com/bayunashr/gopoll/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *gin.Context) {
	tokenString, err := c.Cookie("authorization")
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "error, session expired",
		})
	} else {
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatusJSON(401, gin.H{
					"message": "error, session expired",
				})
			} else {
				var user models.User
				initializers.DB.First(&user, claims["id"])
				if user.ID == 0 {
					c.AbortWithStatusJSON(401, gin.H{
						"message": "error, session expired",
					})
				} else {
					c.Set("user", user)
				}
			}
		} else {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "error, session expired",
			})
		}
		c.Next()
	}
}
