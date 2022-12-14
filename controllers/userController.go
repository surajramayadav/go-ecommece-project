package controllers

import (
	"ecommerce/services"

	"github.com/gin-gonic/gin"
)

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserRegistartion(c)
		return
	}
}

func ViewUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.ViewUserById(c)
		return
	}
}

func ViewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.ViewUser(c)
		return
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UpdateUser(c)
		return
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.DeleteUser(c)
		return
	}
}
