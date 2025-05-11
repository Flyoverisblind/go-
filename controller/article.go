package controller

import (
	"Boke/database"
	"Boke/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Create(&article); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "添加成功", "article": article})
}

func GetArticle(c *gin.Context) {
	id := c.Param("id") // 从 URL 中获取参数，如 /article/1
	var article model.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章未找到"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"article": article})
}
func GetArticles(c *gin.Context) {
	var articles []model.Article
	if err := database.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"articles": articles})

}
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article model.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章未找到"})
		return
	}
	if result := database.DB.Delete(&article); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
