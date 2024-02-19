package model

import (
	"time"
)

type Product struct {
	ID         uint      `json:"id" gorm:"primaryKey, autoIncrement, unique"`
	Name       string    `json:"name"`
	Quantity   int       `json:"quantity"`
	Price      int       `json:"price"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
