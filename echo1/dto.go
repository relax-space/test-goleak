package main

import (
	"time"
)
type Fruit struct {
	Id        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Price     int64     `json:"price"`
	StoreCode string    `json:"storeCode"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}