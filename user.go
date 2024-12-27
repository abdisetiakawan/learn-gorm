package learngorm

import "time"

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