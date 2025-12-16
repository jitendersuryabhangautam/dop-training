package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UsersPostHandlerGin(c *gin.Context) {
	var u User

	// Bind and validate JSON
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON: " + err.Error(),
		})
		return
	}

	// Basic validations
	if u.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	if u.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}
	if u.Age <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "age must be a positive number"})
		return
	}

	id := "usr_" + strconv.Itoa(len(u.Name)+len(u.Email)+u.Age)

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"id":      id,
		"user":    u,
	})
}

func PingHandlerGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func StartServerGin(addr string) error {
	if addr == "" {
		addr = ":8080"
	}

	r := gin.Default()

	r.GET("/ping", PingHandlerGin)
	r.POST("/users", UsersPostHandlerGin)

	println("Server running at http://localhost" + addr + "!")
	return r.Run(addr)
}
