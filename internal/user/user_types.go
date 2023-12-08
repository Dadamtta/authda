package user

import "time"

type userdao struct {
	Id        string    `gorm:"id"`
	HashedPwd string    `gorm:"pwd"`
	Email     string    `gorm:"email"`
	Name      string    `gorm:"user_name"`
	Phone     string    `gorm:"phone"`
	Age       uint8     `gorm:"age"`
	Gender    uint8     `gorm:"gender"`
	Deleted   bool      `gorm:"deleted"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (userdao) TableName() string {
	return "ddt_users"
}
