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

	public := router.Group("/public")
	{
		public.GET("/info", func(ctx *gin.Context) {
			ctx.String(200, "Public information")
		})

		public.GET("/products", func(ctx *gin.Context) {
			ctx.String(200, "Public product list")
		})
	}

	private := router.Group("/private")

	private.Use(AuthMiddleware())
	{
		private.GET("/data", func(ctx *gin.Context) {
			ctx.String(200, "Private data accessible after authentication")
		})
		private.POST("/create", func(ctx *gin.Context) {
			ctx.String(200, "Create a new resource")
		})
	}

	router.Run(":8080")
}
