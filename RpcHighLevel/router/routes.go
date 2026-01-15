package router

import (
	"highlevel/handler"
	"highlevel/proto/product/v1/productv1connect"
	"highlevel/proto/user/v1/userv1connect"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Configure CORS for Connect RPC API
	config := cors.Config{
		AllowOrigins:     []string{"*"}, // In production, specify allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Connect-Protocol-Version", "Connect-Timeout", "Connect-Content-Encoding", "Connect-Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Connect-Content-Encoding", "Connect-Accept-Encoding", "Connect-Protocol-Version", "Connect-Timeout"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// Connect RPC endpoint
	userServiceHandler := handler.NewUserServiceHandler()
	path, connectHandler := userv1connect.NewUserServiceHandler(userServiceHandler)
	r.Any(path+"*path", gin.WrapH(connectHandler))

	// Connect RPC endpoint
	productServiceHandler := handler.NewProductServiceHandler()
	path2, connectHandler2 := productv1connect.NewProductServiceHandler(productServiceHandler)
	r.Any(path2+"*path2", gin.WrapH(connectHandler2))

	return r
}
