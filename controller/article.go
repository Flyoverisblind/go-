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

func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	var article model.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章未找到"})
		return
	}
	// 增加阅读数
	database.DB.Model(&article).Update("Views", article.Views+1)
	c.JSON(http.StatusOK, gin.H{"article": article})
}
func Modification(c *gin.Context) {
	id := c.Param("id") // 从 URL 中获取参数，如 /article/1
	var article model.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章未找到"})
		return
	}
	// 拿到修改的数据并检验
	var updateData model.Article
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据无效"})
		return
	}
	// 更新文章
	if err := database.DB.Model(&article).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新文章失败"})
		return
	}

	// 返回更新后的文章
	c.JSON(http.StatusOK, gin.H{
		"message": "文章更新成功",
		"article": article,
	})

}
