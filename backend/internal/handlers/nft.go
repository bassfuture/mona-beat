package handlers

import (
	"net/http"
	"strconv"

	"nft-capture-game/internal/services"

	"github.com/gin-gonic/gin"
)

type NFTHandler struct {
	nftService *services.NFTService
}

func NewNFTHandler(nftService *services.NFTService) *NFTHandler {
	return &NFTHandler{nftService: nftService}
}

type TransferNFTRequest struct {
	FromWallet string `json:"from_wallet" binding:"required"`
	ToWallet   string `json:"to_wallet" binding:"required"`
}

func (h *NFTHandler) GetUserNFTs(c *gin.Context) {
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

	nfts, err := h.nftService.GetUserNFTs(walletAddress, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"nfts":   nfts,
		"limit":  limit,
		"offset": offset,
	})
}

func (h *NFTHandler) GetNFTByTokenID(c *gin.Context) {
	tokenIDStr := c.Param("tokenId")
	if tokenIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token ID is required"})
		return
	}

	tokenID, err := strconv.ParseUint(tokenIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token ID"})
		return
	}

	nft, err := h.nftService.GetNFTByTokenID(tokenID)
	if err != nil {
		if err.Error() == "NFT not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"nft": nft})
}

func (h *NFTHandler) GetAllNFTs(c *gin.Context) {
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

	nfts, err := h.nftService.GetAllNFTs(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"nfts":   nfts,
		"limit":  limit,
		"offset": offset,
	})
}

func (h *NFTHandler) GetNFTsByRarity(c *gin.Context) {
	rarity := c.Param("rarity")
	if rarity == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rarity is required"})
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

	nfts, err := h.nftService.GetNFTsByRarity(rarity, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"nfts":   nfts,
		"rarity": rarity,
		"limit":  limit,
		"offset": offset,
	})
}

func (h *NFTHandler) GetNFTStats(c *gin.Context) {
	stats, err := h.nftService.GetNFTStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stats": stats})
}

type MintNFTRequest struct {
	CaptureID  string `json:"capture_id" binding:"required"`
	UserID     uint   `json:"user_id" binding:"required"`
	Rarity     string `json:"rarity" binding:"required"`
	Difficulty int    `json:"difficulty" binding:"required,min=1,max=5"`
}

func (h *NFTHandler) MintNFT(c *gin.Context) {
	var req MintNFTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nft, err := h.nftService.MintNFT(req.CaptureID, req.UserID, req.Rarity, req.Difficulty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "NFT minted successfully",
		"nft":     nft,
	})
}

func (h *NFTHandler) GetRecentNFTs(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	nfts, err := h.nftService.GetRecentNFTs(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"nfts":  nfts,
		"limit": limit,
	})
}

func (h *NFTHandler) GetStats(c *gin.Context) {
	stats, err := h.nftService.GetNFTStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stats": stats})
}

func (h *NFTHandler) TransferNFT(c *gin.Context) {
	tokenIDStr := c.Param("tokenId")
	if tokenIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token ID is required"})
		return
	}

	tokenID, err := strconv.ParseUint(tokenIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token ID"})
		return
	}

	var req TransferNFTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.nftService.TransferNFT(tokenID, req.FromWallet, req.ToWallet)
	if err != nil {
		if err.Error() == "NFT not found or not owned by sender" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "sender not found" || err.Error() == "receiver not found" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFT transferred successfully",
		"token_id": tokenID,
		"from": req.FromWallet,
		"to": req.ToWallet,
	})
}