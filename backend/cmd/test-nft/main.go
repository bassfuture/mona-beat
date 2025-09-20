package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"nft-capture-game/pkg/nft"
)

func main() {
	fmt.Println("🎨 测试NFT生成器...")

	// 创建生成器实例
	generator := nft.NewGenerator()
	
	// 设置正确的生成器路径
	generatorPath := filepath.Join("/home/image/nft-capture-game/backend", "nft_generator")
	generator.SetGeneratorPath(generatorPath)

	// 验证生成器设置
	fmt.Println("📋 验证生成器设置...")
	if err := generator.ValidateGenerator(); err != nil {
		log.Fatalf("❌ 生成器验证失败: %v", err)
	}
	fmt.Println("✅ 生成器验证通过")

	// 获取生成器信息
	info := generator.GetGeneratorInfo()
	infoJSON, _ := json.MarshalIndent(info, "", "  ")
	fmt.Printf("📊 生成器信息:\n%s\n\n", infoJSON)

	// 测试生成单个NFT
	fmt.Println("🎲 生成测试NFT...")
	result, err := generator.GenerateSingleNFT(999)
	if err != nil {
		log.Fatalf("❌ NFT生成失败: %v", err)
	}

	fmt.Printf("✅ NFT生成成功!\n")
	fmt.Printf("🧬 DNA: %s\n", result.DNA)
	fmt.Printf("📛 名称: %s\n", result.Metadata.Name)
	fmt.Printf("📝 描述: %s\n", result.Metadata.Description)
	fmt.Printf("🎨 属性数量: %d\n", len(result.Metadata.Attributes))

	// 显示属性
	fmt.Println("\n🏷️  NFT属性:")
	for _, attr := range result.Metadata.Attributes {
		fmt.Printf("  - %s: %s (稀有度: %s)\n", attr.TraitType, attr.Value, attr.Rarity)
	}

	// 测试为捕获生成NFT
	fmt.Println("\n🎯 测试捕获NFT生成...")
	captureResult, err := generator.GenerateForCapture(123, "Dragon")
	if err != nil {
		log.Fatalf("❌ 捕获NFT生成失败: %v", err)
	}

	fmt.Printf("✅ 捕获NFT生成成功!\n")
	fmt.Printf("📛 名称: %s\n", captureResult.Metadata.Name)
	fmt.Printf("📝 描述: %s\n", captureResult.Metadata.Description)

	fmt.Println("\n🎉 所有测试通过!")
}