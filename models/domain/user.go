package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"column:id;primaryKey;type:varchar(100)" json:"id"`
	FullName  string    `json:"full_name " gorm:"column:full_name"`
	Email     string    `gorm:"column:email;unique" json:"email"`
	Password  string    `json:"-" gorm:"column:password"`
	Role      string    `json:"role" gorm:"column:role"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *User) TableName() string  {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		uuidObj := uuid.New()
		u.ID = "user-" + uuidObj.String()
	}
	return
}