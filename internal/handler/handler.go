package handler

import (
	"MedicalDataStorage/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/fillUserInfo", h.FillUserInfo)
		api.GET("/getUserInfo", h.GetUserInfo)
	}
	return router
}
