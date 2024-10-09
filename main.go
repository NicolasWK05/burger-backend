package main

import (
	"net/http"

	"burger.local/internal"
	"burger.local/models"
	"burger.local/pkg"
	"github.com/gin-gonic/gin"
)

func main() {

	var err error

	if gin.EnvGinMode != "release" {
		println("Running development mode")
		err = internal.TestInit()
	} else {
		err = internal.Init()
	}

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/register", func(c *gin.Context) {
		var auth models.AuthInput
		c.BindJSON(&auth)

		// Check if the username is already taken
		user, _ := internal.GetUserByUsername(auth.Username)

		if user != (models.User{}) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
			return
		}

		// Hash the password
		hash, err := pkg.HashPassword(auth.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user = models.User{Username: auth.Username, Password: hash}

		err = internal.InsertUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		token, err := pkg.GenerateToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "User registered"})
	})

	r.POST("/login", func(c *gin.Context) {
		var auth models.AuthInput
		c.BindJSON(&auth)

		user, err := internal.GetUserByUsername(auth.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if user == (models.User{}) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
			return
		}

		if !pkg.ValidatePassword(auth.Password, user.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
			return
		}

		token, err := pkg.GenerateToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
	})

	r.Run()
}
