// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/markcheno/go-gin-app/handlers"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := handlers.InitRoutes()

	router.Run()
}
