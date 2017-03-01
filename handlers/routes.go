package handlers

import "github.com/gin-gonic/gin"

// InitRoutes initialize all routes
func InitRoutes() *gin.Engine {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Use(SetUserStatus())
	router.GET("/", ShowIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", EnsureNotLoggedIn(), ShowLoginPage)
		userRoutes.POST("/login", EnsureNotLoggedIn(), PerformLogin)
		userRoutes.GET("/logout", EnsureLoggedIn(), Logout)
		userRoutes.GET("/register", EnsureNotLoggedIn(), ShowRegistrationPage)
		userRoutes.POST("/register", EnsureNotLoggedIn(), Register)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", GetArticle)
		articleRoutes.GET("/create", EnsureLoggedIn(), ShowArticleCreationPage)
		articleRoutes.POST("/create", EnsureLoggedIn(), CreateArticle)
	}

	return router
}
