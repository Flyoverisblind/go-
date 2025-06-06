package router

import (
	"Boke/controller"
	"Boke/database"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	database.InitMySQL()
	// 文章接口
	r.POST("/AddArticle", controller.AddArticle)
	r.GET("/getArticle/:id", controller.GetArticle)
	r.DELETE("/deleteArticle/:id", controller.DeleteArticle)
	r.GET("/getArticles", controller.GetArticles)
	r.GET("/getArticleByID/:id", controller.GetArticleByID)
	r.POST("/Modification/:id", controller.Modification)
}
