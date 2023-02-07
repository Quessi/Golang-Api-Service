package main

import (
	"gin-project/controller"
	"gin-project/middlewares"
	"gin-project/service"
	"io"
	"os"
	"github.com/gin-gonic/gin"
)

var(
	videoService service.VideoService=service.New()
	videoController controller.VideoController=controller.New(videoService)
)
func setUpLogOutput(){
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	// server := gin.Default()

	// server.GET("/test", func(ctx *gin.Context){
	// 	ctx.JSON(200, gin.H{
	// 		"message":"OK!!",
	// 	})
	// })
	setUpLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(),middlewares.Logger(),middlewares.BasicAuth(),)


	server.GET("/videos", func(ctx *gin.Context){
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos",func(ctx *gin.Context){
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":8080")

}