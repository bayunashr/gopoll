package controllers

import (
	"github.com/bayunashr/gopoll/initializers"
	"github.com/bayunashr/gopoll/models"
	"github.com/gin-gonic/gin"
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
					c.JSON(200, gin.H{
						"message": "success, youre logged in",
						"user":    selUser.Name,
					})
				}
			}
		}
	}
}
