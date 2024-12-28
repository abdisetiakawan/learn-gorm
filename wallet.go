package learngorm

import "time"

type Wallet struct {
	ID        string `gorm:"column:id"`
	UserId    string `gorm:"column:user_id"`
	Balance   int64  `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User	  *User		`gorm:"foreignKey:user_id;references:id"`
}