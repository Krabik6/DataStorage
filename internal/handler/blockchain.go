package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUserInfo(ctx *gin.Context) {
	pass, err := getPassword(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address, err := getAddress(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.services.Queries.GetUserInfo(address, pass)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"info": info})
}

func (h *Handler) FillUserInfo(ctx *gin.Context) {
	pass, err := getPassword(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address, err := getAddress(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := getData(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := h.services.Writes.FillUserInfo(data, pass, address)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tx": tx.Hash().Hex()})
	//сегодня день наобарот ты красишь свои алые губы и выбрав плаьте по ярче ты выйдешь вечером в мир наобарот идешь туда где музыка клубы
}
