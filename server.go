package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-gin-po/controller"
	"gitlab.com/golang-gin-po/middleware"
	"gitlab.com/golang-gin-po/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()
	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middleware.Logger(),
		middleware.BasicAuth())
	//Basic Authorization Middleware applies to "/api" only
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid."})
			}
		})
	}
	// The "/view" endpoints are public (no Authorization required)
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
