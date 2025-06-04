package model

type User struct {
	UserID   uint   `gorm:"primaryKey" json:"user_id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
}

func (User) TableName() string {
	return "oe_user" // 表名是 "user" 而不是默认的 "users"
}
