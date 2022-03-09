package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("points")
	}

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

	return router
}
