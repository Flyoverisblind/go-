package main

import (
	"Boke/controller"
	"Boke/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitMySQL()
	r := gin.Default()
	r.POST("/articles", controller.AddArticle)
	r.GET("/getArticle/:id", controller.GetArticle)
	r.DELETE("/deleteArticle/:id", controller.DeleteArticle)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
