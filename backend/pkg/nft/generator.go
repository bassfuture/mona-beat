package nft

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// NFTAttribute represents an NFT attribute
type NFTAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
	Rarity    string `json:"rarity"`
}

// NFTMetadata represents the complete NFT metadata
type NFTMetadata struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	DNA         string         `json:"dna"`
	Edition     int            `json:"edition"`
	Date        int64          `json:"date"`
	Attributes  []NFTAttribute `json:"attributes"`
	Compiler    string         `json:"compiler"`
}

// GenerationResult contains the result of NFT generation
type GenerationResult struct {
	Metadata *NFTMetadata `json:"metadata"`
	DNA      string       `json:"dna"`
	Success  bool         `json:"success"`
	Error    string       `json:"error,omitempty"`
}

// Generator handles NFT generation using Node.js backend
type Generator struct {
	generatorPath string
	nodeCommand   string
}

// NewGenerator creates a new NFT generator instance
func NewGenerator() *Generator {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		wd = "/home/image/nft-capture-game/backend"
	}
	
	generatorPath := filepath.Join(wd, "nft_generator")
	
	return &Generator{
		generatorPath: generatorPath,
		nodeCommand:   "node",
	}
}

// SetGeneratorPath sets the path to the NFT generator directory
func (g *Generator) SetGeneratorPath(path string) {
	g.generatorPath = path
}

// GenerateSingleNFT generates a single NFT with the given edition number
func (g *Generator) GenerateSingleNFT(editionNumber int) (*GenerationResult, error) {
	// 直接调用单个NFT生成器
	scriptContent := fmt.Sprintf(`
const { generateSingleNFT } = require('./single_nft_generator.js');

(async () => {
  try {
    const result = await generateSingleNFT(%d);
    console.log(JSON.stringify(result));
  } catch (error) {
    console.log(JSON.stringify({
      success: false,
      error: error.message
    }));
  }
})();
`, editionNumber)

	// 写入临时脚本文件
	tempScriptPath := filepath.Join(g.generatorPath, "temp_generate.js")
	err := os.WriteFile(tempScriptPath, []byte(scriptContent), 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to create temp script: %v", err)
	}
	
	// 确保在函数结束时删除临时文件
	defer os.Remove(tempScriptPath)

	// 执行Node.js脚本
	cmd := exec.Command(g.nodeCommand, "temp_generate.js")
	cmd.Dir = g.generatorPath
	
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute NFT generator: %v", err)
	}

	// 解析输出
	var result GenerationResult
	err = json.Unmarshal(output, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse generator output: %v", err)
	}

	if !result.Success {
		return &result, fmt.Errorf("NFT generation failed: %s", result.Error)
	}

	return &result, nil
}

// GenerateForCapture generates an NFT for a successful creature capture
func (g *Generator) GenerateForCapture(userID uint, creatureType string) (*GenerationResult, error) {
	// 使用时间戳和用户ID生成唯一的edition number
	editionNumber := int(time.Now().Unix()%1000000) + int(userID)*1000
	
	result, err := g.GenerateSingleNFT(editionNumber)
	if err != nil {
		return nil, err
	}

	// 自定义NFT名称和描述
	if result.Metadata != nil {
		result.Metadata.Name = fmt.Sprintf("Captured %s #%d", creatureType, editionNumber)
		result.Metadata.Description = fmt.Sprintf("A unique NFT commemorating the successful capture of a %s creature by user %d", creatureType, userID)
	}

	return result, nil
}

// ValidateGenerator checks if the NFT generator is properly set up
func (g *Generator) ValidateGenerator() error {
	// 检查生成器目录是否存在
	if _, err := os.Stat(g.generatorPath); os.IsNotExist(err) {
		return fmt.Errorf("generator directory does not exist: %s", g.generatorPath)
	}

	// 检查simple-generator.js是否存在
	generatorScript := filepath.Join(g.generatorPath, "simple-generator.js")
	if _, err := os.Stat(generatorScript); os.IsNotExist(err) {
		return fmt.Errorf("generator script does not exist: %s", generatorScript)
	}

	// 检查layers目录是否存在
	layersDir := filepath.Join(g.generatorPath, "layers")
	if _, err := os.Stat(layersDir); os.IsNotExist(err) {
		return fmt.Errorf("layers directory does not exist: %s", layersDir)
	}

	// 检查Node.js是否可用
	cmd := exec.Command(g.nodeCommand, "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Node.js is not available: %v", err)
	}

	return nil
}

// GetGeneratorInfo returns information about the generator setup
func (g *Generator) GetGeneratorInfo() map[string]interface{} {
	info := map[string]interface{}{
		"generator_path": g.generatorPath,
		"node_command":   g.nodeCommand,
	}

	// 检查各个组件的状态
	if err := g.ValidateGenerator(); err != nil {
		info["status"] = "error"
		info["error"] = err.Error()
	} else {
		info["status"] = "ready"
	}

	// 获取layers信息
	layersDir := filepath.Join(g.generatorPath, "layers")
	if entries, err := os.ReadDir(layersDir); err == nil {
		layerCount := 0
		for _, entry := range entries {
			if entry.IsDir() {
				layerCount++
			}
		}
		info["layer_count"] = layerCount
	}

	return info
}