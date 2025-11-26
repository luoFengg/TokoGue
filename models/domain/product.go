package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string `gorm:"column:id;primaryKey;type:varchar(100)" json:"id"`
	
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Price       int64  `gorm:"column:price" json:"price"`
	Stock       int    `gorm:"column:stock" json:"stock"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
	}

	func ( p *Product ) TableName() string {
		return "products"
	}

	func ( p *Product ) BeforeCreate(tx *gorm.DB) (err error) {
		if p.ID == "" {
			// Generate UUID baru
			uuidObj := uuid.New()
			p.ID = "product-" + uuidObj.String()
		}
		return
	}