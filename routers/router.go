package routers

import (
	"civi-backend-challenge/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func InitRouter(h handlers.Handler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("points", h.GetPoints)
	}

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

	return router
}
