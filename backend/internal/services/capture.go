package services

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"nft-capture-game/internal/models"

	"gorm.io/gorm"
)

type CaptureService struct {
	db          *gorm.DB
	userService *UserService
	nftService  *NFTService
}

func NewCaptureService(db *gorm.DB, userService *UserService, nftService *NFTService) *CaptureService {
	return &CaptureService{
		db:          db,
		userService: userService,
		nftService:  nftService,
	}
}

func (s *CaptureService) AttemptCapture(walletAddress string, targetType string, difficulty int) (*models.Capture, error) {
	// 获取或创建用户信息
	user, err := s.userService.GetProfile(walletAddress)
	if err != nil {
		if err.Error() == "user not found" {
			// 用户不存在，自动创建用户
			defaultUsername := fmt.Sprintf("User_%s", walletAddress[2:8]) // 使用地址前6位作为默认用户名
			user, err = s.userService.Register(walletAddress, defaultUsername)
			if err != nil {
				return nil, fmt.Errorf("failed to create user: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to get user: %w", err)
		}
	}

	// 计算成功概率（基于难度）
	successRate := s.calculateSuccessRate(difficulty)
	
	// 随机判断是否成功
	isSuccessful := rand.Float64() < successRate

	// 创建捕捉记录
	capture := &models.Capture{
		ID:            fmt.Sprintf("capture_%d_%d", user.ID, time.Now().UnixNano()),
		UserID:        user.ID,
		WalletAddress: walletAddress,
		Success:       isSuccessful,
		Rarity:        targetType,
		Metadata:      fmt.Sprintf(`{"difficulty": %d, "target_type": "%s"}`, difficulty, targetType),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// 开始事务
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 保存捕捉记录
	if err := tx.Create(capture).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create capture: %w", err)
	}

	// 更新用户统计
	if err := s.userService.IncrementCapturesWithTx(tx, walletAddress); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update user captures: %w", err)
	}

	// 如果捕捉成功，创建NFT
	if isSuccessful {
		nft, err := s.nftService.MintNFTWithTx(tx, capture.ID, user.ID, targetType, difficulty)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to mint NFT: %w", err)
		}
		
		capture.NFTTokenID = &nft.TokenID
		if err := tx.Save(capture).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to update capture with NFT: %w", err)
		}

		// 更新用户成功捕捉统计（使用同一事务）
		if err := s.userService.IncrementSuccessfulCapturesWithTx(tx, walletAddress); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to update user successful captures: %w", err)
		}

		// 更新用户NFT统计（使用同一事务）
		if err := s.userService.IncrementNFTsWithTx(tx, walletAddress); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to update user NFTs: %w", err)
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// 重新加载捕捉记录以获取关联数据
	if err := s.db.Preload("User").Where("id = ?", capture.ID).First(capture).Error; err != nil {
		return nil, fmt.Errorf("failed to reload capture: %w", err)
	}
	// 手动加载 NFT（因为 Capture.NFT 使用了 gorm:"-" 不支持关联预加载）
	if capture.NFTTokenID != nil {
		var nft models.NFT
		err := s.db.Where("token_id = ?", *capture.NFTTokenID).First(&nft).Error
		if err == nil {
			capture.NFT = &nft
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to load nft: %w", err)
		}
	}

	return capture, nil
}

func (s *CaptureService) GetUserCaptures(walletAddress string, limit, offset int) ([]models.Capture, error) {
	user, err := s.userService.GetProfile(walletAddress)
	if err != nil {
		if err.Error() == "user not found" {
			// 用户不存在，自动创建用户
			defaultUsername := fmt.Sprintf("User_%s", walletAddress[2:8]) // 使用地址前6位作为默认用户名
			user, err = s.userService.Register(walletAddress, defaultUsername)
			if err != nil {
				return nil, fmt.Errorf("failed to create user: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to get user: %w", err)
		}
	}

	var captures []models.Capture
	err = s.db.Preload("User").
		Where("user_id = ?", user.ID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&captures).Error
	
	if err != nil {
		return nil, fmt.Errorf("failed to get user captures: %w", err)
	}

	// 批量手动加载 NFT
	tokenIDs := make([]uint64, 0, len(captures))
	for _, c := range captures {
		if c.NFTTokenID != nil {
			tokenIDs = append(tokenIDs, *c.NFTTokenID)
		}
	}
	if len(tokenIDs) > 0 {
		var nfts []models.NFT
		if err := s.db.Where("token_id IN ?", tokenIDs).Find(&nfts).Error; err != nil {
			return nil, fmt.Errorf("failed to batch load nfts: %w", err)
		}
		m := make(map[uint64]*models.NFT, len(nfts))
		for i := range nfts {
			n := &nfts[i]
			m[n.TokenID] = n
		}
		for i := range captures {
			if captures[i].NFTTokenID != nil {
				if nftPtr, ok := m[*captures[i].NFTTokenID]; ok {
					captures[i].NFT = nftPtr
				}
			}
		}
	}

	return captures, nil
}

func (s *CaptureService) GetCaptureByID(captureID string) (*models.Capture, error) {
	var capture models.Capture
	err := s.db.Preload("User").Where("id = ?", captureID).First(&capture).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("capture not found")
		}
		return nil, fmt.Errorf("failed to get capture: %w", err)
	}

	if capture.NFTTokenID != nil {
		var nft models.NFT
		if err := s.db.Where("token_id = ?", *capture.NFTTokenID).First(&nft).Error; err == nil {
			capture.NFT = &nft
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to load nft: %w", err)
		}
	}

	return &capture, nil
}

func (s *CaptureService) GetRecentCaptures(limit int) ([]models.Capture, error) {
	var captures []models.Capture
	err := s.db.Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Find(&captures).Error
	
	if err != nil {
		return nil, fmt.Errorf("failed to get recent captures: %w", err)
	}

	// 批量手动加载 NFT
	tokenIDs := make([]uint64, 0, len(captures))
	for _, c := range captures {
		if c.NFTTokenID != nil {
			tokenIDs = append(tokenIDs, *c.NFTTokenID)
		}
	}
	if len(tokenIDs) > 0 {
		var nfts []models.NFT
		if err := s.db.Where("token_id IN ?", tokenIDs).Find(&nfts).Error; err != nil {
			return nil, fmt.Errorf("failed to batch load nfts: %w", err)
		}
		m := make(map[uint64]*models.NFT, len(nfts))
		for i := range nfts {
			n := &nfts[i]
			m[n.TokenID] = n
		}
		for i := range captures {
			if captures[i].NFTTokenID != nil {
				if nftPtr, ok := m[*captures[i].NFTTokenID]; ok {
					captures[i].NFT = nftPtr
				}
			}
		}
	}

	return captures, nil
}

func (s *CaptureService) GetCaptureStats() (map[string]interface{}, error) {
	var totalCaptures int64
	var successfulCaptures int64

	// 获取总捕捉数
	if err := s.db.Model(&models.Capture{}).Count(&totalCaptures).Error; err != nil {
		return nil, fmt.Errorf("failed to get total captures: %w", err)
	}

	// 获取成功捕捉数
	if err := s.db.Model(&models.Capture{}).Where("success = ?", true).Count(&successfulCaptures).Error; err != nil {
		return nil, fmt.Errorf("failed to get successful captures: %w", err)
	}

	successRate := float64(0)
	if totalCaptures > 0 {
		successRate = float64(successfulCaptures) / float64(totalCaptures) * 100
	}

	return map[string]interface{}{
		"total_captures":      totalCaptures,
		"successful_captures": successfulCaptures,
		"success_rate":        successRate,
	}, nil
}

// calculateSuccessRate 根据难度计算成功率
func (s *CaptureService) calculateSuccessRate(difficulty int) float64 {
	// 基础成功率配置
	baseRates := map[int]float64{
		1: 0.8,  // 简单 - 80%
		2: 0.6,  // 中等 - 60%
		3: 0.4,  // 困难 - 40%
		4: 0.2,  // 极难 - 20%
		5: 0.1,  // 传说 - 10%
	}

	if rate, exists := baseRates[difficulty]; exists {
		return rate
	}
	
	// 默认中等难度
	return 0.5
}