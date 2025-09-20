package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"nft-capture-game/internal/models"
	"nft-capture-game/internal/services"
	"nft-capture-game/pkg/blockchain"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 设置环境变量（如果没有设置的话）
	if os.Getenv("MYSQL_HOST") == "" {
		os.Setenv("MYSQL_HOST", "localhost")
	}
	if os.Getenv("MYSQL_PORT") == "" {
		os.Setenv("MYSQL_PORT", "3306")
	}
	if os.Getenv("MYSQL_USER") == "" {
		os.Setenv("MYSQL_USER", "root")
	}
	if os.Getenv("MYSQL_PASSWORD") == "" {
		os.Setenv("MYSQL_PASSWORD", "911make@")
	}
	if os.Getenv("MYSQL_DATABASE") == "" {
		os.Setenv("MYSQL_DATABASE", "nft_backend")
	}

	// 创建数据库配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移数据库表 - 按依赖关系顺序创建
	// 1. 先创建User表（被其他表引用）
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("User表迁移失败: %v", err)
	}
	
	// 2. 创建Capture表（引用User表）
	err = db.AutoMigrate(&models.Capture{})
	if err != nil {
		log.Fatalf("Capture表迁移失败: %v", err)
	}
	
	// 3. 最后创建NFT表（引用User表）
	err = db.AutoMigrate(&models.NFT{})
	if err != nil {
		log.Fatalf("NFT表迁移失败: %v", err)
	}

	// 创建NFT服务
	nftService := services.NewNFTService(db)

	// 启动API服务器
	fmt.Println("\n启动API服务器...")
	startAPIServer(db, nftService)
}

func generateNFT(db *gorm.DB, nftService *services.NFTService) (*models.NFT, error) {

	// 创建测试用户
	user := &models.User{
		WalletAddress: "0x1234567890123456789012345678901234567890",
		Username:      "demo_user",
		Email:         "demo@example.com",
	}

	// 先检查用户是否存在，如果不存在则创建
	var existingUser models.User
	result := db.Where("wallet_address = ?", user.WalletAddress).First(&existingUser)
	if result.Error != nil {
		// 用户不存在，创建新用户
		if err := db.Create(user).Error; err != nil {
			return nil, fmt.Errorf("创建用户失败: %v", err)
		}
	} else {
		// 用户已存在，使用现有用户
		user = &existingUser
	}

	// 创建测试捕获记录
	capture := &models.Capture{
		ID:            "demo_capture_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		UserID:        user.ID,
		WalletAddress: user.WalletAddress,
		Success:       true,
		Rarity:        "common",
		Metadata:      `{"type": "demo", "location": "test_area", "creature_type": "demo_creature"}`,
	}

	// 创建捕获记录
	if err := db.Create(capture).Error; err != nil {
		return nil, fmt.Errorf("创建捕获记录失败: %v", err)
	}

	// 生成NFT
	return nftService.MintNFT(capture.ID, user.ID, "common", 1)
}

func startAPIServer(db *gorm.DB, nftService *services.NFTService) {
	// CORS中间件
	corsMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next(w, r)
		}
	}

	// 设置静态文件服务
	http.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir("../../storage"))))

	// 提供图片文件服务
	http.HandleFunc("/image/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		imagePath := r.URL.Path[7:] // 去掉 "/image/" 前缀
		fullPath := filepath.Join("/home/image/nft-capture-game/backend/nft_generator/build/images", imagePath)
		
		// 检查文件是否存在
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		
		http.ServeFile(w, r, fullPath)
	}))

	// API路由
	// 1. 生成NFT接口
	http.HandleFunc("/api/nft/generate", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// 解析请求参数
		var req struct {
			UserID   uint   `json:"user_id"`
			Rarity   string `json:"rarity"`
			Quantity int    `json:"quantity"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// 设置默认值
		if req.UserID == 0 {
			req.UserID = 1
		}
		if req.Rarity == "" {
			req.Rarity = "common"
		}
		if req.Quantity == 0 {
			req.Quantity = 1
		}

		// 创建用户（如果不存在）
		user := &models.User{
			Username: fmt.Sprintf("user_%d", req.UserID),
			Email:    fmt.Sprintf("user%d@example.com", req.UserID),
		}
		db.FirstOrCreate(user, models.User{Username: user.Username})

		// 创建捕获记录
		capture := &models.Capture{
			ID:            "api_capture_" + strconv.FormatInt(time.Now().UnixNano(), 10),
			UserID:        user.ID,
			WalletAddress: user.WalletAddress,
			Success:       true,
			Rarity:        req.Rarity,
			Metadata:      `{"type": "api", "source": "frontend", "location": "API_Generated"}`,
		}

		if err := db.Create(capture).Error; err != nil {
			http.Error(w, "Failed to create capture record", http.StatusInternalServerError)
			return
		}

		// 生成NFT
		nft, err := nftService.MintNFT(capture.ID, user.ID, req.Rarity, req.Quantity)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate NFT: %v", err), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"nft":     nft,
			"message": "NFT generated successfully",
		})
	}))

	// 2. 获取NFT详情接口
	http.HandleFunc("/api/nft/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// 从URL路径中提取NFT ID
		path := r.URL.Path
		if len(path) <= 9 { // "/api/nft/" 长度为9
			http.Error(w, "NFT ID required", http.StatusBadRequest)
			return
		}

		nftIDStr := path[9:] // 去掉 "/api/nft/" 前缀
		nftID, err := strconv.ParseUint(nftIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid NFT ID", http.StatusBadRequest)
			return
		}

		var nft models.NFT
		if err := db.First(&nft, nftID).Error; err != nil {
			http.Error(w, "NFT not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(nft)
	}))

	// 3. 获取所有NFT列表接口
	http.HandleFunc("/api/nfts", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		var nfts []models.NFT
		if err := db.Find(&nfts).Error; err != nil {
			http.Error(w, "Failed to fetch NFTs", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"nfts":    nfts,
			"count":   len(nfts),
		})
	}))

	// 4. 健康检查接口
	http.HandleFunc("/api/health", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    "ok",
			"timestamp": time.Now(),
			"service":   "NFT API Server",
		})
	}))

	// 创建用户服务
	userService := services.NewUserService(db)

	// 6. 开始游戏会话接口
	http.HandleFunc("/api/game/start", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		var req struct {
			WalletAddress string  `json:"wallet_address"`
			PaymentAmount float64 `json:"payment_amount"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.WalletAddress == "" {
			http.Error(w, "Wallet address is required", http.StatusBadRequest)
			return
		}

		user, err := userService.StartGameSession(req.WalletAddress, req.PaymentAmount)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to start game session: %v", err), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":        true,
			"message":        "Game session started successfully",
			"remaining_taps": user.RemainingTaps,
			"game_active":    user.GameActive,
			"user_id":        user.ID,
		})
	}))

	// 7. 敲击接口
	http.HandleFunc("/api/game/tap", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		var req struct {
			WalletAddress string `json:"wallet_address"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.WalletAddress == "" {
			http.Error(w, "Wallet address is required", http.StatusBadRequest)
			return
		}

		result, err := userService.ProcessTap(req.WalletAddress)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to process tap: %v", err), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(result)
	}))

	// 8. 获取游戏状态接口
	http.HandleFunc("/api/game/status", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		walletAddress := r.URL.Query().Get("wallet_address")
		if walletAddress == "" {
			http.Error(w, "Wallet address is required", http.StatusBadRequest)
			return
		}

		user, err := userService.GetByWalletAddress(walletAddress)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":        true,
			"remaining_taps": user.RemainingTaps,
			"game_active":    user.GameActive,
			"total_captures": user.TotalCaptures,
			"total_nfts":     user.TotalNFTs,
			"session_start":  user.SessionStartAt,
			"last_tap":       user.LastTapAt,
		})
	}))

	// 9. 区块链铸造NFT接口
	http.HandleFunc("/api/blockchain/mint", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		var req struct {
			PlayerAddress string `json:"player_address"`
			CaptureID     string `json:"capture_id"`
			Difficulty    uint64 `json:"difficulty"`
			Rarity        uint8  `json:"rarity"`
			Success       bool   `json:"success"`
			TokenURI      string `json:"token_uri"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.PlayerAddress == "" {
			http.Error(w, "Player address is required", http.StatusBadRequest)
			return
		}

		// 验证钱包地址格式
		if !common.IsHexAddress(req.PlayerAddress) {
			http.Error(w, "Invalid wallet address format", http.StatusBadRequest)
			return
		}

		// 验证地址长度（以太坊地址应该是42字符，包含0x前缀）
		if len(req.PlayerAddress) != 42 {
			http.Error(w, "Invalid wallet address length", http.StatusBadRequest)
			return
		}

		// 设置默认值
		if req.CaptureID == "" {
			req.CaptureID = fmt.Sprintf("capture_%d", time.Now().UnixNano())
		}
		if req.Difficulty == 0 {
			req.Difficulty = 1
		}
		if req.TokenURI == "" {
			req.TokenURI = "https://example.com/metadata/1.json"
		}

		// 创建区块链客户端
		config := blockchain.GetDefaultConfig()
		client, err := blockchain.NewBlockchainClient(config)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create blockchain client: %v", err), http.StatusInternalServerError)
			return
		}
		defer client.Close()

		// 调用智能合约铸造NFT
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		
		playerAddr := common.HexToAddress(req.PlayerAddress)
		
		// 更新Gas价格
		if err := client.UpdateGasPrice(ctx); err != nil {
			log.Printf("Warning: Failed to update gas price: %v", err)
		}

		// 调用合约的attemptCapture函数
		tx, err := client.GetContract().AttemptCapture(
			client.GetAuth(),
			playerAddr,
			req.CaptureID,
			big.NewInt(int64(req.Difficulty)),
			req.Rarity,
			req.Success,
			req.TokenURI,
		)
		if err != nil {
			// 详细的错误分类
			errorMsg := fmt.Sprintf("Failed to attempt capture: %v", err)
			if strings.Contains(err.Error(), "insufficient funds") {
				errorMsg = "Insufficient funds for gas fees"
			} else if strings.Contains(err.Error(), "nonce") {
				errorMsg = "Transaction nonce error, please try again"
			} else if strings.Contains(err.Error(), "gas") {
				errorMsg = "Gas estimation failed or gas limit too low"
			}
			
			http.Error(w, errorMsg, http.StatusInternalServerError)
			return
		}

		// 等待交易被挖矿（可选，用于验证）
		var txStatus string
		var blockNumber uint64
		
		// 尝试等待交易确认，但不阻塞太久
		go func() {
			select {
			case <-ctx.Done():
				return
			default:
				// 异步等待交易确认
				if receipt, err := client.GetClient().TransactionReceipt(ctx, tx.Hash()); err == nil {
					if receipt.Status == 1 {
						log.Printf("Transaction %s confirmed in block %d", tx.Hash().Hex(), receipt.BlockNumber.Uint64())
					} else {
						log.Printf("Transaction %s failed in block %d", tx.Hash().Hex(), receipt.BlockNumber.Uint64())
					}
				}
			}
		}()
		
		txStatus = "pending"
		blockNumber = 0

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":            true,
			"transaction_hash":   tx.Hash().Hex(),
			"transaction_status": txStatus,
			"block_number":       blockNumber,
			"message":            "NFT capture attempt submitted successfully",
			"player_address":     req.PlayerAddress,
			"capture_id":         req.CaptureID,
			"token_uri":          req.TokenURI,
			"gas_used":           tx.Gas(),
			"gas_price":          tx.GasPrice().String(),
		})
	}))

	// 5. API文档接口
	http.HandleFunc("/api", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		apiDoc := map[string]interface{}{
			"name":        "NFT Capture Game API",
			"version":     "1.0.0",
			"description": "API for NFT generation and management",
			"endpoints": map[string]interface{}{
				"POST /api/nft/generate":     "Generate a new NFT",
				"GET /api/nft/{id}":          "Get NFT details by ID",
				"GET /api/nfts":              "Get all NFTs",
				"GET /api/health":            "Health check",
				"POST /api/game/start":       "Start a new game session (100 taps for payment)",
				"POST /api/game/tap":         "Process a tap (chance to get NFT)",
				"GET /api/game/status":       "Get current game status",
				"POST /api/blockchain/mint":  "Mint NFT on blockchain (requires player_address)",
				"GET /api":                   "API documentation",
			},
		}
		json.NewEncoder(w).Encode(apiDoc)
	}))

	fmt.Println("API服务器启动在 http://localhost:8080")
	fmt.Println("本地访问: http://localhost:8080")
	fmt.Println("外部访问: http://59.110.161.193:8080")
	fmt.Println("API文档: http://localhost:8080/api")
	
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}