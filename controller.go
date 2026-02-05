package main

import (
		"github.com/gin-gonic/gin"

)

func createRouter(router *gin.RouterGroup){
	router.GET("/hello-world", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message" : "hello-world !",
		})
	})
}