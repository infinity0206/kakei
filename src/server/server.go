package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"main/controller"
)

// 初期化
func Init() {
	r := router()

	r.Run()
}

// ルーティング
func router() *gin.Engine {
	r := gin.Default()

	// CORS対応
	r.Use(CORS())

	// ルーティング
	u := r.Group("/assignees")
	{
		ctrl := controller.AssigneeController{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
	}

	u = r.Group("/tasks")
	{
		ctrl := controller.TaskController{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
	}

	return r
}

// CORS
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
