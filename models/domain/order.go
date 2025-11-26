package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID         string    `json:"id" gorm:"primaryKey;column:id;type:varchar(100)"`
	UserID     string    `json:"user_id" gorm:"column:user_id;type:varchar(100);not null"`
	Status     string    `json:"status" gorm:"column:status;type:varchar(20)"`
	TotalPrice int64     `json:"total_price" gorm:"column:total_price"`
	PaymentURL string	`json:"payment_url" gorm:"column:payment_url;type:text"`
	
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	// Relasi: 1 order punya banyak items
	Items []OrderItem `json:"items" gorm:"foreignKey:OrderID"`

	// Relasi: 1 Order milik 1 User
	User User `json:"-" gorm:"foreignKey:UserID"`
	
}

type OrderItem struct {
	ID		string `json:"id" gorm:"primaryKey;column:id;type:varchar(100)"`
	OrderID string `json:"order_id" gorm:"column:order_id;type:varchar(100);not null"`
	ProductID string `json:"product_id" gorm:"column:product_id;type:varchar(100);not null"`
	Quantity  int    `json:"quantity" gorm:"column:quantity"`
	Price     int  `json:"price" gorm:"column:price"` // harga satuan pas beli

	// Relasi: item punya info Product
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}

func (order *Order) TableName() string  {
	return "orders"
}

func (orderItem *OrderItem) TableName() string  {
	return "order_items"
}

func ( o *Order ) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == "" {
		// Generate UUID baru
		uuidObj := uuid.New()
		o.ID = "order-" + uuidObj.String()
	}
	return
}

func ( oi *OrderItem ) BeforeCreate(tx *gorm.DB) (err error) {
	if oi.ID == "" {
		// Generate UUID baru
		uuidObj := uuid.New()
		oi.ID = "orderitem-" + uuidObj.String()
	}
	return
}