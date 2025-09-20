package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"nft-capture-game/pkg/nft"
)

func main() {
	fmt.Println("🗄️  测试NFT元数据存储功能...")

	// 创建NFT生成器
	generator := nft.NewGenerator()
	generator.SetGeneratorPath("/home/image/nft-capture-game/backend/nft_generator")

	// 验证生成器
	fmt.Println("📋 验证生成器设置...")
	if err := generator.ValidateGenerator(); err != nil {
		log.Fatal("❌ 生成器验证失败:", err)
	}
	fmt.Println("✅ 生成器验证通过")

	// 生成NFT
	fmt.Println("🎨 生成NFT...")
	result, err := generator.GenerateForCapture(123, "Dragon")
	if err != nil {
		log.Fatal("❌ NFT生成失败:", err)
	}

	if !result.Success {
		log.Fatal("❌ NFT生成失败:", result.Error)
	}

	fmt.Println("✅ NFT生成成功!")
	fmt.Printf("📛 名称: %s\n", result.Metadata.Name)
	fmt.Printf("📝 描述: %s\n", result.Metadata.Description)
	fmt.Printf("🧬 DNA: %s\n", result.DNA)
	fmt.Printf("🖼️  图片: %s\n", result.Metadata.Image)

	// 测试元数据保存功能
	fmt.Println("\n💾 测试元数据保存...")
	
	// 创建存储目录
	storageDir := "/home/image/nft-capture-game/backend/storage/nft_metadata"
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		log.Fatal("❌ 创建存储目录失败:", err)
	}

	// 保存元数据到文件
	metadataFile := fmt.Sprintf("%s/test_nft_%d.json", storageDir, result.Metadata.Edition)
	
	// 将元数据转换为JSON并保存
	import "encoding/json"
	metadataJSON, err := json.MarshalIndent(result.Metadata, "", "  ")
	if err != nil {
		log.Fatal("❌ 序列化元数据失败:", err)
	}

	if err := os.WriteFile(metadataFile, metadataJSON, 0644); err != nil {
		log.Fatal("❌ 保存元数据文件失败:", err)
	}

	fmt.Printf("✅ 元数据已保存到: %s\n", metadataFile)

	// 验证文件内容
	savedData, err := os.ReadFile(metadataFile)
	if err != nil {
		log.Fatal("❌ 读取保存的元数据失败:", err)
	}

	fmt.Printf("📄 保存的元数据大小: %d 字节\n", len(savedData))
	
	// 显示属性信息
	fmt.Printf("🏷️  属性数量: %d\n", len(result.Metadata.Attributes))
	for i, attr := range result.Metadata.Attributes {
		if i < 3 { // 只显示前3个属性
			fmt.Printf("   - %s: %s (稀有度: %s)\n", attr.TraitType, attr.Value, attr.Rarity)
		}
	}

	fmt.Println("\n🎉 NFT元数据存储测试完成!")
}