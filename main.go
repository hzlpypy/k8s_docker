package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	gin.SetMode("debug")
	engine := gin.New()
	engine.Use(logService)
	engine.GET("/api/v1/test_docker/task", func(c *gin.Context) {
		 c.JSON(http.StatusOK, gin.H{
		 	"err": "",
		 	"msg": "ok",
		 })
	})
	err := engine.Run(":8035")
	if err != nil{
		log.Fatal(err)
	}
}

func logService(c *gin.Context) {
	log.Println(fmt.Sprintf("access:%d",c.Request.ContentLength))
	c.Next()
}

