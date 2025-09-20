package handlers

import (
	"net/http"
	"strconv"

	"nft-capture-game/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

type RegisterRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	Username      string `json:"username" binding:"required"`
}

type UpdateProfileRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type StartGameRequest struct {
	WalletAddress string  `json:"wallet_address" binding:"required"`
	PaymentAmount float64 `json:"payment_amount" binding:"required"`
}

type TapRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
}

type GameStatusResponse struct {
	RemainingTaps int    `json:"remaining_taps"`
	GameActive    bool   `json:"game_active"`
	Message       string `json:"message"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Register(req.WalletAddress, req.Username)
	if err != nil {
		if err.Error() == "user already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	walletAddress := c.Param("wallet")
	if walletAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wallet address is required"})
		return
	}

	user, err := h.userService.GetProfile(walletAddress)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	walletAddress := c.Param("wallet")
	if walletAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wallet address is required"})
		return
	}

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.UpdateProfile(walletAddress, req.Username, req.Email)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user":    user,
	})
}

func (h *UserHandler) GetStats(c *gin.Context) {
	walletAddress := c.Param("wallet")
	if walletAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wallet address is required"})
		return
	}

	user, err := h.userService.GetStats(walletAddress)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stats": user})
}

func (h *UserHandler) GetLeaderboard(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	users, err := h.userService.GetLeaderboard(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"leaderboard": users,
		"limit":       limit,
	})
}

// StartGame 开始游戏会话 - 用户付费获得100次敲击机会
func (h *UserHandler) StartGame(c *gin.Context) {
	var req StartGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 站内账本模式：允许 payment_amount 作为充值；如果小于0仍不合法
	if req.PaymentAmount < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment amount"})
		return
	}

	user, err := h.userService.StartGameSession(req.WalletAddress, req.PaymentAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Game session started successfully",
		"remaining_taps": user.RemainingTaps,
		"game_active":    user.GameActive,
		"user":           user,
	})
}

// Tap 处理用户敲击
func (h *UserHandler) Tap(c *gin.Context) {
	var req TapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.ProcessTap(req.WalletAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetGameStatus 获取用户游戏状态
func (h *UserHandler) GetGameStatus(c *gin.Context) {
	walletAddress := c.Param("wallet_address")
	if walletAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wallet address is required"})
		return
	}

	user, err := h.userService.GetByWalletAddress(walletAddress)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, GameStatusResponse{
		RemainingTaps: user.RemainingTaps,
		GameActive:    user.GameActive,
		Message:       "Game status retrieved successfully",
	})
}