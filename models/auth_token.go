package models

import "time"

type AuthToken struct{
	ID 			int64 			`json:"id" gorm:"column:id"`
	UserId 	int64				`json:"userId" gorm:"column:userId"`
	Token 	string			`json:"token" gorm:"column:token"`

	CreatedAt 	time.Time		`json:"createdAt" gorm:"column:created_at"`
	UpdatedAt 	time.Time		`json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt 	time.Time		`json:"deletedAt" gorm:"column:deleted_at"`
}