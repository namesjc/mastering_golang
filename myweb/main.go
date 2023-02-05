package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/form", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "form.html", gin.H{})
	})

	// router.GET("/data", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "data.html", gin.H{"data": "Hello Go/Gin World."})
	// })

	// router.GET("/json", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"result":    "ok",
	// 		"data":      "Hello Go/Gin World.",
	// 		"developer": "echo",
	// 	})
	// })

	router.POST("/service", func(ctx *gin.Context) {
		uname := ctx.PostForm("uname")
		ctx.JSON(http.StatusOK, gin.H{
			"result": "ok",
			"hello":  uname,
		})
	})
	router.Run(":8000")
}
