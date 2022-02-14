package Routes

import (
	Controller "key-rate-api/src/Controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/keyrate", Controller.KeyRate)
	}

	return router
}
