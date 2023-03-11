package handler

import (
	"Mehmat/Api/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("./templates/**/*")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	program := router.Group("/task")
	{
		program.GET("/send", h.Program)
		program.POST("/send", h.Compile)
	}

	return router
}
