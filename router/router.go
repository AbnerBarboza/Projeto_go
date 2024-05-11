package router

import "github.com/gin-gonic/gin"

func Initialize(){
		//inibilizando o gin
		r := gin.Default()
		//definindo uma rota
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run(":8000")
}