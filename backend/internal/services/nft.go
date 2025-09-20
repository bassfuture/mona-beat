package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"nft-capture-game/internal/models"
	"nft-capture-game/pkg/nft"

	"gorm.io/gorm"
)

type NFTService struct {
	db           *gorm.DB
	nftGenerator *nft.Generator
}

func NewNFTService(db *gorm.DB) *NFTService {
	// 初始化NFT生成器
	generator := nft.NewGenerator()
	generator.SetGeneratorPath("/home/image/nft-capture-game/backend/nft_generator")
	
	return &NFTService{
		db:           db,
		nftGenerator: generator,
	}
}

func (s *NFTService) MintNFT(captureID string, userID uint, rarity string, difficulty int) (*models.NFT, error) {
	// 默认使用全局DB（非事务）
	return s.MintNFTWithTx(s.db, captureID, userID, rarity, difficulty)
}

func (s *NFTService) MintNFTWithTx(tx *gorm.DB, captureID string, userID uint, rarity string, difficulty int) (*models.NFT, error) {
	// 生成唯一的TokenID（这里简化处理，实际应该与区块链交互）
	tokenID := uint64(time.Now().UnixNano())

	// 使用NFT生成器为捕捉生成NFT
	generationResult, err := s.nftGenerator.GenerateForCapture(userID, rarity)
	if err != nil {
		return nil, fmt.Errorf("failed to generate NFT: %w", err)
	}

	// 确保生成成功
	if !generationResult.Success || generationResult.Metadata == nil {
		return nil, fmt.Errorf("NFT generation failed: %s", generationResult.Error)
	}

	metadata := generationResult.Metadata

	// 保存完整的元数据到文件系统
	metadataPath, fullMetadataJSON, err := s.saveMetadataToFile(tokenID, metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to save metadata: %w", err)
	}

	nft := &models.NFT{
		TokenID:         tokenID,
		UserID:          userID,
		CaptureID:       captureID,
		Name:            metadata.Name,
		Description:     metadata.Description,
		ImageURL:        metadata.Image,
		MetadataURL:     fmt.Sprintf("/metadata/%d.json", tokenID),
		Rarity:          rarity,
		Attributes:      s.convertAttributesToJSON(metadata.Attributes),
		// 新增字段
		FullMetadata:    fullMetadataJSON,
		DNA:             generationResult.DNA,
		ImagePath:       metadata.Image, // 目前使用相对路径
		MetadataPath:    metadataPath,
		TransactionHash: "", // 实际应该从区块链获取
		Minted:          false,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := tx.Create(nft).Error; err != nil {
		return nil, fmt.Errorf("failed to create NFT: %w", err)
	}

	return nft, nil
}

func (s *NFTService) GetUserNFTs(walletAddress string, limit, offset int) ([]models.NFT, error) {
	var nfts []models.NFT
	err := s.db.Joins("JOIN users ON users.id = nfts.user_id").
		Where("users.wallet_address = ?", walletAddress).
		Preload("User").
		Preload("Capture").
		Order("nfts.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&nfts).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get user NFTs: %w", err)
	}

	return nfts, nil
}

func (s *NFTService) GetNFTByTokenID(tokenID uint64) (*models.NFT, error) {
	var nft models.NFT
	err := s.db.Preload("User").Preload("Capture").Where("token_id = ?", tokenID).First(&nft).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("NFT not found")
		}
		return nil, fmt.Errorf("failed to get NFT: %w", err)
	}

	return &nft, nil
}

func (s *NFTService) GetAllNFTs(limit, offset int) ([]models.NFT, error) {
	var nfts []models.NFT
	err := s.db.Preload("User").
		Preload("Capture").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&nfts).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get NFTs: %w", err)
	}

	return nfts, nil
}

func (s *NFTService) GetNFTsByRarity(rarity string, limit, offset int) ([]models.NFT, error) {
	var nfts []models.NFT
	err := s.db.Where("rarity = ?", rarity).
		Preload("User").
		Preload("Capture").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&nfts).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get NFTs by rarity: %w", err)
	}

	return nfts, nil
}

func (s *NFTService) GetNFTStats() (map[string]interface{}, error) {
	var totalNFTs int64
	if err := s.db.Model(&models.NFT{}).Count(&totalNFTs).Error; err != nil {
		return nil, fmt.Errorf("failed to get total NFTs: %w", err)
	}

	// 按稀有度统计
	rarityStats := make(map[string]int64)
	rarities := []string{"common", "uncommon", "rare", "epic", "legendary"}
	
	for _, rarity := range rarities {
		var count int64
		if err := s.db.Model(&models.NFT{}).Where("rarity = ?", rarity).Count(&count).Error; err != nil {
			return nil, fmt.Errorf("failed to get %s NFT count: %w", rarity, err)
		}
		rarityStats[rarity] = count
	}

	return map[string]interface{}{
		"total_nfts":     totalNFTs,
		"rarity_stats":   rarityStats,
	}, nil
}

func (s *NFTService) GetRecentNFTs(limit int) ([]models.NFT, error) {
	var nfts []models.NFT
	err := s.db.Preload("User").
		Preload("Capture").
		Order("nfts.created_at DESC").
		Limit(limit).
		Find(&nfts).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get recent NFTs: %w", err)
	}

	return nfts, nil
}

func (s *NFTService) TransferNFT(tokenID uint64, fromWallet, toWallet string) error {
	// 获取发送方用户
	var fromUser models.User
	if err := s.db.Where("wallet_address = ?", fromWallet).First(&fromUser).Error; err != nil {
		return fmt.Errorf("sender not found: %w", err)
	}

	// 获取接收方用户
	var toUser models.User
	if err := s.db.Where("wallet_address = ?", toWallet).First(&toUser).Error; err != nil {
		return fmt.Errorf("receiver not found: %w", err)
	}

	// 获取NFT
	var nft models.NFT
	if err := s.db.Where("token_id = ? AND user_id = ?", tokenID, fromUser.ID).First(&nft).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("NFT not found or not owned by sender")
		}
		return fmt.Errorf("failed to get NFT: %w", err)
	}

	// 转移NFT
	nft.UserID = toUser.ID
	nft.UpdatedAt = time.Now()

	if err := s.db.Save(&nft).Error; err != nil {
		return fmt.Errorf("failed to transfer NFT: %w", err)
	}

	return nil
}

// generateMetadata 生成NFT元数据
func (s *NFTService) generateMetadata(rarity string, difficulty int) string {
	// 这里简化处理，实际应该生成完整的JSON元数据
	return fmt.Sprintf(`{"rarity": "%s", "difficulty": %d, "created_at": "%s"}`, 
		rarity, difficulty, time.Now().Format(time.RFC3339))
}

// convertAttributesToJSON 将NFT属性转换为JSON字符串
func (s *NFTService) convertAttributesToJSON(attributes []nft.NFTAttribute) string {
	if len(attributes) == 0 {
		return "[]"
	}
	
	jsonData, err := json.Marshal(attributes)
	if err != nil {
		// 如果序列化失败，返回空数组
		return "[]"
	}
	
	return string(jsonData)
}

// saveMetadataToFile 保存元数据到文件系统
func (s *NFTService) saveMetadataToFile(tokenID uint64, metadata *nft.NFTMetadata) (string, string, error) {
	// 创建存储目录
	storageDir := "/home/image/nft-capture-game/backend/storage/nft_metadata"
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		return "", "", fmt.Errorf("failed to create storage directory: %w", err)
	}

	// 生成文件路径
	filename := fmt.Sprintf("nft_%d.json", tokenID)
	filePath := filepath.Join(storageDir, filename)

	// 序列化元数据
	metadataJSON, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal metadata: %w", err)
	}

	// 写入文件
	if err := os.WriteFile(filePath, metadataJSON, 0644); err != nil {
		return "", "", fmt.Errorf("failed to write metadata file: %w", err)
	}

	return filePath, string(metadataJSON), nil
}