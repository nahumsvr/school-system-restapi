package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	SubjectID uint   `gorm:"primaryKey;autoIncrement" json:"subject_id"`
	Name      string `gorm:"type:varchar(100);not null" json:"name"`
}
