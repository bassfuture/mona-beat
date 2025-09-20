package handlers

import (
	"net/http"
	"strconv"

	"nft-capture-game/internal/services"

	"github.com/gin-gonic/gin"
)

type CaptureHandler struct {
	captureService *services.CaptureService
}

func NewCaptureHandler(captureService *services.CaptureService) *CaptureHandler {
	return &CaptureHandler{captureService: captureService}
}

type AttemptCaptureRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	TargetType    string `json:"target_type" binding:"required"`
	Difficulty    int    `json:"difficulty" binding:"required,min=1,max=5"`
}

func (h *CaptureHandler) AttemptCapture(c *gin.Context) {
	var req AttemptCaptureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	capture, err := h.captureService.AttemptCapture(req.WalletAddress, req.TargetType, req.Difficulty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Capture attempt completed",
		"capture": capture,
	})
}

func (h *CaptureHandler) GetUserCaptures(c *gin.Context) {
	walletAddress := c.Param("wallet")
	if walletAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wallet address is required"})
		return
	}

	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	captures, err := h.captureService.GetUserCaptures(walletAddress, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"captures": captures,
		"limit":    limit,
		"offset":   offset,
	})
}

func (h *CaptureHandler) GetCaptureByID(c *gin.Context) {
	captureID := c.Param("id")
	if captureID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capture ID is required"})
		return
	}

	capture, err := h.captureService.GetCaptureByID(captureID)
	if err != nil {
		if err.Error() == "capture not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"capture": capture})
}

func (h *CaptureHandler) GetRecentCaptures(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	captures, err := h.captureService.GetRecentCaptures(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"captures": captures,
		"limit":    limit,
	})
}

func (h *CaptureHandler) GetCaptureStats(c *gin.Context) {
	stats, err := h.captureService.GetCaptureStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stats": stats})
}