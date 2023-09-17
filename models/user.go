package models

import "time"

type User struct{
	ID					uint64		`json:"id" gorm:"primaryKey;column:id"`
	UUID 				string		`json:"uuid" gorm:"column:uuid"`
	FullName 		string 		`json:"fullName" gorm:"column:fullName"`
	Email 			string		`json:"email" gorm:"column:email"`
	Password 		string 		`json:"password" gorm:"type:text;column:password"`
	Role 				int				`json:"role" gorm:"column:role"`

	CreatedAt		time.Time	`json:"createdAt" gorm:"column:created_at"`
	UpdatedAt		time.Time	`json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt		time.Time	`json:"deletedAt" gorm:"column:deleted_at"`
}