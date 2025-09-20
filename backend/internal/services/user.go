package services

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"nft-capture-game/internal/models"
	"nft-capture-game/pkg/blockchain"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Register(walletAddress, username string) (*models.User, error) {
	// 检查用户是否已存在
	var existingUser models.User
	err := s.db.Where("wallet_address = ?", walletAddress).First(&existingUser).Error
	if err == nil {
		return nil, errors.New("user already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	// 创建新用户
	user := &models.User{
		WalletAddress:      walletAddress,
		Username:           username,
		TotalCaptures:      0,
		SuccessfulCaptures: 0,
		TotalNFTs:          0,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetProfile(walletAddress string) (*models.User, error) {
	var user models.User
	err := s.db.Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}

	return &user, nil
}

func (s *UserService) UpdateProfile(walletAddress, username, email string) (*models.User, error) {
	var user models.User
	err := s.db.Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// 更新用户信息
	if username != "" {
		user.Username = username
	}
	if email != "" {
		user.Email = email
	}
	user.UpdatedAt = time.Now()

	if err := s.db.Save(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &user, nil
}

func (s *UserService) GetStats(walletAddress string) (*models.User, error) {
	var user models.User
	err := s.db.Preload("Captures").Preload("NFTs").Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user stats: %w", err)
	}

	return &user, nil
}

func (s *UserService) GetLeaderboard(limit int) ([]models.User, error) {
	var users []models.User
	err := s.db.Order("successful_captures DESC, total_captures DESC").Limit(limit).Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get leaderboard: %w", err)
	}

	return users, nil
}

func (s *UserService) IncrementCaptures(walletAddress string) error {
	return s.db.Model(&models.User{}).
		Where("wallet_address = ?", walletAddress).
		Updates(map[string]interface{}{
			"total_captures":  gorm.Expr("total_captures + 1"),
			"last_capture_at": time.Now(),
			"updated_at":      time.Now(),
		}).Error
}

func (s *UserService) IncrementCapturesWithTx(tx *gorm.DB, walletAddress string) error {
	return tx.Model(&models.User{}).
		Where("wallet_address = ?", walletAddress).
		Updates(map[string]interface{}{
			"total_captures":  gorm.Expr("total_captures + 1"),
			"last_capture_at": time.Now(),
			"updated_at":      time.Now(),
		}).Error
}

func (s *UserService) IncrementSuccessfulCaptures(walletAddress string) error {
	return s.db.Model(&models.User{}).
		Where("wallet_address = ?", walletAddress).
		Updates(map[string]interface{}{
			"successful_captures": gorm.Expr("successful_captures + 1"),
			"updated_at":          time.Now(),
		}).Error
}

func (s *UserService) IncrementSuccessfulCapturesWithTx(tx *gorm.DB, walletAddress string) error {
	return tx.Model(&models.User{}).
		Where("wallet_address = ?", walletAddress).
		Updates(map[string]interface{}{
			"successful_captures": gorm.Expr("successful_captures + 1"),
			"updated_at":          time.Now(),
		}).Error
}

func (s *UserService) IncrementNFTs(walletAddress string) error {
	return s.db.Model(&models.User{}).
		Where("wallet_address = ?", walletAddress).
		Updates(map[string]interface{}{
			"total_nf_ts": gorm.Expr("total_nf_ts + 1"),
			"updated_at": time.Now(),
		}).Error
}

func (s *UserService) IncrementNFTsWithTx(tx *gorm.DB, walletAddress string) error {
	return tx.Model(&models.User{}).
		Where("wallet_address = ?", walletAddress).
		Updates(map[string]interface{}{
			"total_nf_ts": gorm.Expr("total_nf_ts + 1"),
			"updated_at": time.Now(),
		}).Error
}

// GetByWalletAddress 根据钱包地址获取用户
func (s *UserService) GetByWalletAddress(walletAddress string) (*models.User, error) {
	var user models.User
	err := s.db.Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

// StartGameSession 开始游戏会话 - 用户付费获得100次敲击机会
func (s *UserService) StartGameSession(walletAddress string, paymentAmount float64) (*models.User, error) {
	// 获取或创建用户
	user, err := s.GetByWalletAddress(walletAddress)
	if err != nil {
		// 如果用户不存在，创建新用户
		if err.Error() == "user not found" {
			user = &models.User{
				WalletAddress: walletAddress,
				Username:      fmt.Sprintf("User_%s", walletAddress[2:8]), // 使用地址前6位作为默认用户名
			}
			if err := s.db.Create(user).Error; err != nil {
				return nil, fmt.Errorf("failed to create user: %w", err)
			}
		} else {
			return nil, err
		}
	}

	// 检查是否已有活跃游戏会话
	if user.GameActive && user.RemainingTaps > 0 {
		return nil, errors.New("user already has an active game session")
	}

	// 开始新的游戏会话
	now := time.Now()
	user.RemainingTaps = 100 // 固定100次敲击机会
	user.GameActive = true
	user.SessionStartAt = &now
	user.UpdatedAt = now

	if err := s.db.Save(user).Error; err != nil {
		return nil, fmt.Errorf("failed to start game session: %w", err)
	}

	return user, nil
}

// TapResult 敲击结果结构
type TapResult struct {
	Success       bool        `json:"success"`
	RemainingTaps int         `json:"remaining_taps"`
	GameActive    bool        `json:"game_active"`
	NFTGenerated  bool        `json:"nft_generated"`
	NFT           *models.NFT `json:"nft,omitempty"`
	Message       string      `json:"message"`
}

// ProcessTap 处理用户敲击
func (s *UserService) ProcessTap(walletAddress string) (*TapResult, error) {
	user, err := s.GetByWalletAddress(walletAddress)
	if err != nil {
		return nil, err
	}

	// 检查游戏状态
	if !user.GameActive {
		return nil, errors.New("no active game session")
	}

	if user.RemainingTaps <= 0 {
		return nil, errors.New("no remaining taps")
	}

	// 减少敲击次数
	user.RemainingTaps--
	now := time.Now()
	user.LastTapAt = &now
	user.UpdatedAt = now

	// 如果敲击次数用完，结束游戏
	if user.RemainingTaps <= 0 {
		user.GameActive = false
	}

	// 计算NFT生成概率 (这里设置为10%的概率)
	nftGenerated := false
	var nft *models.NFT
	var txHash string
	
	// 使用当前时间作为随机种子
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < 0.1 { // 10% 概率
		nftGenerated = true
		
		// 自动调用区块链铸造NFT
		txHash, err = s.mintNFTForUser(walletAddress)
		if err != nil {
			// 如果区块链铸造失败，记录错误但不影响游戏流程
			fmt.Printf("Failed to mint NFT on blockchain for user %s: %v\n", walletAddress, err)
			// 仍然标记为生成了NFT，但没有区块链交易
		}
		
		// 创建NFT记录
		tokenID := uint64(time.Now().Unix()) // 使用时间戳作为TokenID
		nft = &models.NFT{
			UserID:        user.ID,
			TokenID:       tokenID,
			WalletAddress: walletAddress,
			CaptureID:     fmt.Sprintf("capture_%s_%d", walletAddress[2:8], time.Now().Unix()),
			Name:          fmt.Sprintf("Capture NFT #%d", user.TotalNFTs+1),
			Description:   "A unique NFT captured through the NFT Capture Game",
			ImageURL:      "https://example.com/nft-image.png", // 这里应该是实际的图片URL
			MetadataURL:   "https://example.com/metadata.json", // 这里应该是实际的元数据URL
			Rarity:        "common",
			Attributes:    "{}",  // 使用有效的空JSON对象而不是空字符串
			FullMetadata:  "{}",  // 使用有效的空JSON对象
			TransactionHash: txHash,
			Minted:        txHash != "",
			CreatedAt:     now,
		}
		
		// 保存NFT到数据库
		if err := s.db.Create(nft).Error; err != nil {
			fmt.Printf("Failed to save NFT to database: %v\n", err)
		}
	}

	// 更新用户数据
	if err := s.db.Save(user).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// 更新统计
	s.IncrementCaptures(walletAddress)
	if nftGenerated {
		s.IncrementSuccessfulCaptures(walletAddress)
		s.IncrementNFTs(walletAddress)
	}

	result := &TapResult{
		Success:       true,
		RemainingTaps: user.RemainingTaps,
		GameActive:    user.GameActive,
		NFTGenerated:  nftGenerated,
		NFT:           nft,
		Message:       fmt.Sprintf("Tap processed. %d taps remaining.", user.RemainingTaps),
	}

	if nftGenerated {
		if txHash != "" {
			result.Message = fmt.Sprintf("Congratulations! You got an NFT! Transaction: %s", txHash)
		} else {
			result.Message = "Congratulations! You got an NFT! (Blockchain minting failed)"
		}
	}

	if user.RemainingTaps <= 0 {
		result.Message = "Game session completed!"
	}

	return result, nil
}

// mintNFTForUser 为用户铸造NFT到区块链
func (s *UserService) mintNFTForUser(walletAddress string) (string, error) {
	// 创建区块链客户端
	config := blockchain.GetDefaultConfig()
	client, err := blockchain.NewBlockchainClient(config)
	if err != nil {
		return "", fmt.Errorf("failed to create blockchain client: %w", err)
	}
	defer client.Close()

	// 验证钱包地址格式
	if !common.IsHexAddress(walletAddress) {
		return "", fmt.Errorf("invalid wallet address format: %s", walletAddress)
	}

	// 生成唯一的capture_id
	captureID := fmt.Sprintf("capture_%s_%d", walletAddress[2:8], time.Now().Unix())

	// 创建上下文，设置30秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	playerAddr := common.HexToAddress(walletAddress)
	
	// 更新Gas价格
	if err := client.UpdateGasPrice(ctx); err != nil {
		fmt.Printf("Warning: Failed to update gas price: %v\n", err)
	}

	// 调用合约的attemptCapture函数
	tx, err := client.GetContract().AttemptCapture(
		client.GetAuth(),
		playerAddr,
		captureID,
		big.NewInt(1), // difficulty
		uint8(1),      // rarity (1 = common)
		true,          // success
		"https://example.com/metadata.json", // tokenURI
	)
	if err != nil {
		return "", fmt.Errorf("failed to attempt capture: %w", err)
	}

	return tx.Hash().Hex(), nil
}