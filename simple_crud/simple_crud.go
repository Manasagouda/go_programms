package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/users", CreateUser)
	r.GET("/users/:id", GetUserByID)
	r.GET("/users", GetAllUsers)
	r.Run()
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "John"},
	{ID: 2, Name: "Jane"},
	{ID: 3, Name: "Bob"},
}

func CreateUser(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	uId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	for _, user := range users {
		if user.ID == uId {
			ctx.JSON(http.StatusOK, user)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
