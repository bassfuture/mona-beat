package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"nft-capture-game/internal/config"
	"nft-capture-game/internal/database"
	"nft-capture-game/internal/models"
	"nft-capture-game/internal/services"
)

func main() {
	fmt.Println("=== MySQL NFT存储测试 ===")

	// 设置环境变量使用MySQL
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "911make@")
	os.Setenv("DB_NAME", "nft_capture_game")

	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	fmt.Printf("连接MySQL数据库: %s:%s\n", cfg.Database.Host, cfg.Database.Port)

	// 连接数据库
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移
	err = database.Migrate(db)
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	fmt.Println("✓ 数据库连接和迁移成功")

	// 创建测试用户
	user := &models.User{
		WalletAddress: "0x1234567890123456789012345678901234567890",
		Username:      "test_user",
	}
	result := db.Create(user)
	if result.Error != nil {
		log.Fatalf("创建用户失败: %v", result.Error)
	}

	fmt.Printf("✓ 创建测试用户: %s\n", user.Username)

	// 创建测试捕捉记录
	capture := &models.Capture{
		ID:            "test_capture_001",
		UserID:        user.ID,
		WalletAddress: user.WalletAddress,
		Success:       true,
		Rarity:        "common",
		Metadata:      `{"test": "data"}`,
	}
	result = db.Create(capture)
	if result.Error != nil {
		log.Fatalf("创建捕捉记录失败: %v", result.Error)
	}

	fmt.Printf("✓ 创建测试捕捉记录: %s\n", capture.ID)

	// 创建NFT服务
	nftService := services.NewNFTService(db)

	// 生成NFT
	nftData, err := nftService.MintNFT(capture.ID, user.ID, "common", 1)
	if err != nil {
		log.Fatalf("生成NFT失败: %v", err)
	}

	fmt.Printf("✓ NFT生成成功:\n")
	fmt.Printf("  Token ID: %d\n", nftData.TokenID)
	fmt.Printf("  DNA: %s\n", nftData.DNA)
	fmt.Printf("  图片路径: %s\n", nftData.ImagePath)
	fmt.Printf("  元数据路径: %s\n", nftData.MetadataPath)

	// 验证文件是否存在
	if _, err := os.Stat(nftData.ImagePath); os.IsNotExist(err) {
		log.Printf("警告: 图片文件不存在: %s", nftData.ImagePath)
	} else {
		fmt.Printf("✓ 图片文件存在: %s\n", nftData.ImagePath)
	}

	if _, err := os.Stat(nftData.MetadataPath); os.IsNotExist(err) {
		log.Printf("警告: 元数据文件不存在: %s", nftData.MetadataPath)
	} else {
		fmt.Printf("✓ 元数据文件存在: %s\n", nftData.MetadataPath)
	}

	// 验证数据库中的数据
	var savedNFT models.NFT
	result = db.First(&savedNFT, nftData.TokenID)
	if result.Error != nil {
		log.Fatalf("查询NFT失败: %v", result.Error)
	}

	fmt.Printf("✓ 数据库验证:\n")
	fmt.Printf("  Token ID: %d\n", savedNFT.TokenID)
	fmt.Printf("  DNA: %s\n", savedNFT.DNA)
	
	// 解析并显示完整元数据
	if savedNFT.FullMetadata != "" {
		var metadata map[string]interface{}
		err = json.Unmarshal([]byte(savedNFT.FullMetadata), &metadata)
		if err == nil {
			fmt.Printf("  完整元数据: %+v\n", metadata)
		}
	}

	fmt.Println("\n=== MySQL NFT存储测试完成 ===")
}