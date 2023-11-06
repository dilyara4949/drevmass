package domain

import "time"



type User struct {
	ID        uint   `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Email     string `gorm:"unique; not null"`
	Password  string `gorm:" not null"`
	Name      string `gorm:" not null"`
	Gender    uint 
	Height    uint 
	Weight    uint 
	Birth     time.Time
	Activity  uint 
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (User, error)
	GetByID( id uint) (User, error)
}
