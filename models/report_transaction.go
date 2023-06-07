package models

import "time"

type ReportTransaction struct{
	ID					uint64		`json:"id" gorm:"primaryKey;column:"`
	UUID				int64			`json:"UUID" gorm:"column:uuid"`
	UserId			int64			`json:"userId" gorm:"column:userId"`
	CategoryId	int64			`json:"categoryId" gorm:"column:categoryId"`
	Nominal 		int64			`json:"nominal" gorm:"column:nominal"`
	
	CreatedAt		time.Time	`json:"createdAt" gorm:"column:created_at"`
	UpdatedAt		time.Time	`json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt		time.Time	`json:"deletedAt" gorm:"column:deleted_at"`
}