package models

type User struct {
	UserID   uint64 `gorm:"primary_key:auto_increment" json:"user_id"`
	Name     string `gorm:"type:varchar(50)" json:"name"`
	Email    string `gorm:"type:varchar(50)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"type:text" json:"token,omitempty"`
}
