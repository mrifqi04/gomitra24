package models

type User struct {
	UserID   uint64 `gorm:"primary_key:auto_increment" json:"user_id"`
	Name     string `gorm:"type:varchar(50)" json:"name"`
	Email    string `gorm:"type:varchar(50)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Phone    string `gorm:"type:varchar(12)" json:"phone"`
	Address  string `gorm:"type:text" json:"address"`
	Token    string `gorm:"type:text" json:"-"`
}
