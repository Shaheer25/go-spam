package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Shaheer25/go-jwt/initializers"
	"github.com/Shaheer25/go-jwt/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//get the email and password
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to add Email and Password",
		})

		return
	}

	//Hash the Password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash the Passoword",
		})
		return
	}

	//create the Email and Password
	user := model.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create the User",
		})
		return
	}

	//Respond

	c.JSON(http.StatusOK, gin.H{})

}
func Login(c *gin.Context) {

	//take user Email and Passoword

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to add Email and Password",
		})

		return
	}

	//Look up requested User
	var user model.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email and Password Didn't match",
		})
		return
	}

	//Check whethre Email and Pass Matched

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Password and Email Didn't match",
			})
		}
		return
	}

	//Generate jwt Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Token",
		})
	}
	//Showcase

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",tokenString,3600 * 24 * 30 *12 ,"","",false,true)

	c.JSON(http.StatusOK, gin.H{})

}


func Validate(c *gin.Context){

	user , _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message":user,
	})
}

