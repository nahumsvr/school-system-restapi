package services

import (
	"fmt"

	"github.com/nahumsvr/school-system-restapi/db"
	"github.com/nahumsvr/school-system-restapi/models"
)

type GradeService struct {
	grades []models.Grade
}

func NewGradeService() *GradeService {
	return &GradeService{
		grades: []models.Grade{},
	}
}

func (s *GradeService) GetAll() []models.Grade {
	grades := db.DB.Find(&s.grades)
	err := grades.Error
	if err != nil {
		fmt.Println("Error fetching grades:", err)
		return nil
	}
	return s.grades
}

func (s *GradeService) Create(grade models.Grade) models.Grade {
	fmt.Println("Creating grade:", grade)
	createdGrade := db.DB.Create(&grade)
	err := createdGrade.Error
	if err != nil {
		fmt.Println("Error creating grade:", err)
		return models.Grade{}
	}
	// Preload Student and Subject after creation
	var result models.Grade
	db.DB.Preload("Student").Preload("Subject").First(&result, grade.GradeID)
	return result
}

func (s *GradeService) Get(id int) (grade models.Grade, err error) {
	db.DB.Preload("Student").Preload("Subject").First(&grade, id)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error fetching grade:", err)
		return models.Grade{}, err
	}
	return grade, nil
}

func (s *GradeService) Update(id int, updatedGrade models.Grade) (grade models.Grade, err error) {
	var g models.Grade
	db.DB.First(&g, id)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error fetching grade:", err)
		return models.Grade{}, err
	}
	g.Grade = updatedGrade.Grade
	g.StudentID = updatedGrade.StudentID
	g.SubjectID = updatedGrade.SubjectID
	db.DB.Save(&g)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error updating grade:", err)
		return models.Grade{}, err
	}
	return g, nil
}

func (s *GradeService) Delete(id int) error {
	var grade models.Grade
	db.DB.Delete(&grade, id)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error deleting grade:", err)
		return err
	}
	return nil
}
