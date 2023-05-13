package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	FirstName string         `json:"first_name" gorm:"type:char(50);not null"`
	LastName  string         `json:"last_name" gorm:"type:char(50);not null"`
	Email     string         `json:"email" gorm:"type:char(50);not null"`
	Phone     string         `json:"phone" gorm:"type:char(30);not null"`
	Course    *Course        `gorm:"-"`
	CreatedAt string         `json:"-"`
	UpdatedAt string         `json:"-"`
	Deleted   gorm.DeletedAt `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}