package router

import (
	"Goprojects/Practice/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(useCase usecase.UserInterface) gin.Engine {
	router := gin.Default()
	router.GET("healthcheck",
		func(c *gin.Context) {
			c.JSON(200, gin.H{"healthcheck": "ok"})
		})
	v1:= router.Group("/api/v1"){
		NewHandler(v1, useCase)
	}
	return router
}
