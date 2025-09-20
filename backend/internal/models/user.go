package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID                 uint           `json:"id" gorm:"primaryKey"`
	WalletAddress      string         `json:"wallet_address" gorm:"uniqueIndex;not null;size:42"`
	Username           string         `json:"username" gorm:"size:50"`
	Email              string         `json:"email" gorm:"size:100"`
	Balance            float64        `json:"balance" gorm:"default:0"`
	TotalCaptures      int            `json:"total_captures" gorm:"default:0"`
	SuccessfulCaptures int            `json:"successful_captures" gorm:"default:0"`
	TotalNFTs          int            `json:"total_nfts" gorm:"column:total_nf_ts;default:0"`
	LastCaptureAt      *time.Time     `json:"last_capture_at"`
	
	// 游戏会话相关字段
	RemainingTaps      int            `json:"remaining_taps" gorm:"default:0"`        // 剩余敲击次数
	GameActive         bool           `json:"game_active" gorm:"default:false"`       // 游戏是否激活
	SessionStartAt     *time.Time     `json:"session_start_at"`                       // 游戏会话开始时间
	LastTapAt          *time.Time     `json:"last_tap_at"`                           // 最后一次敲击时间
	
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联
	Captures []Capture `json:"captures,omitempty" gorm:"foreignKey:UserID"`
	NFTs     []NFT     `json:"nfts,omitempty" gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}