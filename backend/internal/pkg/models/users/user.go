package users

import (
	"time"
	"users_api/internal/pkg/models"

	"gorm.io/gorm"
)

type User struct {
	models.Model
	Username  string `gorm:"column:username;not null;unique_index:username" json:"username" form:"username"`
	Firstname string `gorm:"column:firstname;not null;" json:"firstname" form:"firstname"`
	Lastname  string `gorm:"column:lastname;not null;" json:"lastname" form:"lastname"`
	Hash      string `gorm:"column:hash;not null;" json:"-"`
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *User) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
