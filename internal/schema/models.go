package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string  `json:"name" gorm:"not null"`
	Amount int     `json:"amount" gorm:"not null"`
	Orders []Order `json:"orders"`
}

type Order struct {
	gorm.Model
	UserLineUid string `json:"userLineUid" gorm:"not null"`
	Status      string `json:"status" gorm:"not null"`

	ProductID uint    `json:"productId"`
	Product   Product `json:"product"`
}

type OrderMessage struct {
	ProductName string `json:"productName"`
	UserID      string `json:"userId"`
	OrderID     uint   `json:"orderId"`
}

type LineMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type LineBody struct {
	To       string        `json:"to"`
	Messages []LineMessage `json:"messages"`
}