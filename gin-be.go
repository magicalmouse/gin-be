package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", ctx.Request.Method, ctx.Writer.Status(), duration)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("API-Key")
		if apiKey == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		ctx.Next()
	}
}
func main() {
	router := gin.Default()

	//router.Use(LoggerMiddleware())
	authGroup := router.Group("/api")

	authGroup.Use(AuthMiddleware())
	{
		authGroup.GET("/data", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Authenticated and authrized!"})
		})
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello World!")
	})

	router.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(200, "User ID: "+id)
	})

	router.GET("/search", func(ctx *gin.Context) {
		query := ctx.DefaultQuery("q", "Hey Nice to meet you")
		ctx.String(200, "Search query: "+query)
	})

	router.Run(":8080")
}
