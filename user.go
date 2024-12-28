package learngorm

import (
	"time"

	"gorm.io/gorm"
)

// field gorm ini sebenarnya tidak digunakan tidak masalah, karna sudah otomatis dari go
type User struct {
	ID        string	`gorm:"column:id;<-:create"` 
	Name  	  Name		`gorm:"embedded"`
	Password  string	`gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string 	`gorm:"-"`
}

func (u *User) TableName() string {
	return "users"
}

type Name struct {
	FirstName 	string
	MiddleName 	string
	LastName	string
}

type UserLog struct {
	ID        int		`gorm:"column:id;autoIncrement"` 
	UserId    string	`gorm:"column:user_id"`
	Action    string	`gorm:"column:action"`
	CreatedAt int64 	`gorm:"column:created_at;autoCreateTime:milli;<-:create"`
	UpdatedAt int64 	`gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

type Todo struct {
	gorm.Model
	UserId    		string				`gorm:"column:user_id"`
	Title     		string				`gorm:"column:title"`
	Description     string				`gorm:"column:description"`
}