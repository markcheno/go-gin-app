package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/markcheno/go-gin-app/models"
)

// ShowIndexPage render home page
func ShowIndexPage(c *gin.Context) {

	articles := models.GetAllArticles()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

// ShowArticleCreationPage render create new article page
func ShowArticleCreationPage(c *gin.Context) {

	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

// GetArticle render article
func GetArticle(c *gin.Context) {

	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {

		if article, err := models.GetArticleByID(articleID); err == nil {

			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {

			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {

		c.AbortWithStatus(http.StatusNotFound)
	}
}

// CreateArticle render create article page
func CreateArticle(c *gin.Context) {

	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := models.CreateNewArticle(title, content); err == nil {

		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {

		c.AbortWithStatus(http.StatusBadRequest)
	}
}
