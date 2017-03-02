package main

import (
	"github.com/gin-gonic/gin"
	m "github.com/markcheno/go-gin-app/middleware"
	"github.com/markcheno/go-gin-app/models"
	"github.com/markcheno/go-gin-app/views"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	//db, err := models.NewPostgresDB("host=localhost port=5432 user=postgres dbname=mark sslmode=disable")
	db, err := models.NewSqliteDB("prod.db")
	if err != nil {
		panic(err)
	}

	userView := &views.UserView{DB: db}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Use(m.SetUserStatus())
	router.GET("/", views.ShowIndexPage)

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/login", m.EnsureNotLoggedIn(), userView.ShowLoginPage)
		userRoutes.POST("/login", m.EnsureNotLoggedIn(), userView.PerformLogin)
		userRoutes.GET("/logout", m.EnsureLoggedIn(), userView.Logout)
		userRoutes.GET("/register", m.EnsureNotLoggedIn(), userView.ShowRegistrationPage)
		userRoutes.POST("/register", m.EnsureNotLoggedIn(), userView.Register)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", views.GetArticle)
		articleRoutes.GET("/create", m.EnsureLoggedIn(), views.ShowArticleCreationPage)
		articleRoutes.POST("/create", m.EnsureLoggedIn(), views.CreateArticle)
	}

	router.Run()
}
