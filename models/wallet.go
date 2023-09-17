package models

import "time"

type Wallet struct{
	ID			uint64		`json:"id" gorm:"primaryKey;column:id"`
	Title 	string		`json:"title" gorm:"column:title"`
	UserId 	int64			`json:"userId" gorm:"column:userId"`
	Nominal int64			`json:"nominal" gorm:"column:nominal"`

	CreatedAt		time.Time	`json:"createdAt" gorm:"column:created_at"`
	UpdatedAt		time.Time	`json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt		time.Time	`json:"deletedAt" gorm:"column:deleted_at"`
}