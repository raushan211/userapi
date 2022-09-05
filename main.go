package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	UserID   string `json:"user_id"`
	Mobile   string `json:"mobile"`
	Mail     string `json:"mail"`
	City     string `json:"ciyy"`
	Password string `json:"password" binding:"required"`
}

var Data map[string]User

func main() {
	Data = make(map[string]User)
	r := gin.Default()
	setupRoutes(r)
	r.Run()
}

func setupRoutes(r *gin.Engine) {
	r.POST("/user", CreateUser)
}

func CreateUser(c *gin.Context) {
	reqBody := User{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID == "" {
		res := gin.H{
			"error": "UserId must not be empty",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	Data[reqBody.UserID] = reqBody
	res := gin.H{
		"success": true,
		"user":    reqBody,
	}
	c.JSON(http.StatusOK, res)
	return
}
