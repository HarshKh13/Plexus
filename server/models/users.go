package models

type Users struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email; not null"`
}
