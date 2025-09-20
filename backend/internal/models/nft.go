package models

import (
	"time"
	"gorm.io/gorm"
)

type NFT struct {
	TokenID         uint64         `json:"token_id" gorm:"primaryKey"`
	UserID          uint           `json:"user_id" gorm:"not null;index"`
	WalletAddress   string         `json:"wallet_address" gorm:"not null;index;size:42"`
	CaptureID       string         `json:"capture_id" gorm:"not null;index;size:64"`
	Name            string         `json:"name" gorm:"not null;size:100"`
	Description     string         `json:"description" gorm:"type:text"`
	ImageURL        string         `json:"image_url" gorm:"not null;size:500"`
	MetadataURL     string         `json:"metadata_url" gorm:"not null;size:500"`
	Rarity          string         `json:"rarity" gorm:"not null;size:20"`
	Attributes      string         `json:"attributes" gorm:"type:json"`
	// 新增字段：存储完整的元数据JSON
	FullMetadata    string         `json:"full_metadata" gorm:"type:json"`
	// 新增字段：存储DNA信息
	DNA             string         `json:"dna" gorm:"size:500"`
	// 新增字段：存储图片文件路径（本地存储）
	ImagePath       string         `json:"image_path" gorm:"size:500"`
	// 新增字段：存储元数据文件路径（本地存储）
	MetadataPath    string         `json:"metadata_path" gorm:"size:500"`
	TransactionHash string         `json:"transaction_hash" gorm:"size:66"`
	Minted          bool           `json:"minted" gorm:"default:false"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联
	User    User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Capture Capture `json:"capture,omitempty" gorm:"foreignKey:CaptureID;references:ID"`
}

func (NFT) TableName() string {
	return "nfts"
}