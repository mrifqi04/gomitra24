package models

type Product struct {
	ProductID   uint64 `gorm:"primary_key:auto_increment" json:"product_id"`
	ProductName string `gorm:"type:varchar(50)" json:"product_name"`
	Description string `gorm:"type:text" json:"description"`
	Price       uint32 `gorm:"-" json:"token,omitempty"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
