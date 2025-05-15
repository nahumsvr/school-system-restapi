package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	StudentID uint   `gorm:"primaryKey;autoIncrement" json:"student_id"`
	Name      string `gorm:"type:varchar(100);not null" json:"name"`
	Group     string `gorm:"type:varchar(50);not null" json:"group"`
	Email     string `gorm:"type:varchar(100);unique;not null" json:"email"`
}

// type StudentResponse struct {
// 	Name      string `json:"name"`
// 	Email     string `json:"email"`
// 	StudentID string `json:"student_id"`
// 	Group     string `json:"group"`
// }
