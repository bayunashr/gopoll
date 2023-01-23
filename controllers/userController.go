package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/bayunashr/gopoll/initializers"
	"github.com/bayunashr/gopoll/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user struct {
		Email    string
		Password string
		Name     string
	}
	if c.Bind(&user) != nil {
		c.JSON(400, gin.H{
			"message": "error, fail to read body",
		})
	} else {
		if user.Email == "" || user.Password == "" || user.Name == "" {
			c.JSON(400, gin.H{
				"message": "error, please fill all the required fields",
			})
		} else {
			hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			if err != nil {
				c.JSON(400, gin.H{
					"message": "error, fail to encrypt password",
				})
			} else {
				newUser := models.User{Email: user.Email, Password: string(hashedPass), Name: user.Name}
				result := initializers.DB.Create(&newUser)
				if result.Error != nil {
					c.JSON(400, gin.H{
						"message": "error, fail to create new user",
					})
				} else {
					c.JSON(200, gin.H{
						"message": "success, created new user",
					})
				}
			}
		}
	}
}

func LogIn(c *gin.Context) {
	var user struct {
		Email    string
		Password string
	}
	if c.Bind(&user) != nil {
		c.JSON(400, gin.H{
			"message": "error, fail to read body",
		})
	} else {
		if user.Email == "" || user.Password == "" {
			c.JSON(400, gin.H{
				"message": "error, please fill all the required fields",
			})
		} else {
			var selUser models.User
			initializers.DB.Where("email = ?", user.Email).First(&selUser)
			if selUser.ID == 0 {
				c.JSON(400, gin.H{
					"message": "error, email not found",
				})
			} else {
				err := bcrypt.CompareHashAndPassword([]byte(selUser.Password), []byte(user.Password))
				if err != nil {
					c.JSON(400, gin.H{
						"message": "error, wrong password",
					})
				} else {
					token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
						"id":  selUser.ID,
						"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
					})
					tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
					if err != nil {
						c.JSON(400, gin.H{
							"message": "error, fail to create token",
						})
					} else {
						c.SetSameSite(http.SameSiteLaxMode)
						c.SetCookie("authorization", tokenString, 3600*24*30, "", "", false, true)
						c.JSON(200, gin.H{
							"message": "success, youre logged in",
						})
					}
				}
			}
		}
	}
}

func LogOut(c *gin.Context) {
	c.SetCookie("authorization", "", -1, "", "", false, true)
	c.JSON(200, gin.H{
		"message": "success, youre logged out",
	})
}
