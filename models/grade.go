package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	GradeID   uint    `gorm:"primaryKey;autoIncrement" json:"grade_id"`
	StudentID uint    `gorm:"not null" json:"student_id"`
	SubjectID uint    `gorm:"not null" json:"subject_id"`
	Grade     float64 `gorm:"type:numeric(4,2);check:grade >= 0 AND grade <= 100" json:"grade"`

	Student Student `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE" json:"student"`
	Subject Subject `gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE" json:"subject"`
}
