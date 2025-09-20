package models

import (
	"time"
	"gorm.io/gorm"
)

type Capture struct {
	ID              string         `json:"id" gorm:"primaryKey;size:64"`
	UserID          uint           `json:"user_id" gorm:"not null;index"`
	WalletAddress   string         `json:"wallet_address" gorm:"not null;index;size:42"`
	Success         bool           `json:"success" gorm:"not null"`
	Rarity          string         `json:"rarity" gorm:"size:20"`
	NFTTokenID      *uint64        `json:"nft_token_id" gorm:"index"`
	TransactionHash string         `json:"transaction_hash" gorm:"size:66"`
	Metadata        string         `json:"metadata" gorm:"type:json"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联 - 移除外键约束，避免循环依赖
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	NFT  *NFT `json:"nft,omitempty" gorm:"-"`
}

func (Capture) TableName() string {
	return "captures"
}