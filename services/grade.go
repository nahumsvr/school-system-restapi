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

func (s *GradeService) GetAll() (gradesR []models.GradeResponse) {
	grades := db.DB.Find(&s.grades)
	err := grades.Error
	if err != nil {
		fmt.Println("Error fetching grades:", err)
		return nil
	}
	for _, grade := range s.grades {
		gradesR = append(gradesR, models.GradeResponse{
			GradeID:   grade.GradeID,
			StudentID: grade.StudentID,
			SubjectID: grade.SubjectID,
			Grade:     grade.Grade,
		})
	}
	return gradesR
}

func (s *GradeService) Create(grade *models.Grade) {
	if createdGrade := db.DB.Create(&grade); createdGrade.Error != nil {
		fmt.Println("Error creating grade:", createdGrade.Error)
	}

	StudentService := &StudentService{}
	student, studentErr := StudentService.Get(int(grade.StudentID))
	if studentErr != nil {
		fmt.Println("Error fetching student:", studentErr)
	}
	SubjectService := &SubjectService{}
	subject, subjectErr := SubjectService.Get(int(grade.SubjectID))
	if subjectErr != nil {
		fmt.Println("Error fetching subject:", subjectErr)
	}

	grade.Student = student
	grade.Subject = subject
}

func (s *GradeService) Get(id int) (grade models.GradeResponse, err error) {
	var g models.Grade
	db.DB.Preload("Student").Preload("Subject").First(&g, id)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error fetching grade:", err)
		return models.GradeResponse{}, err
	}
	grade.GradeID = g.GradeID
	grade.StudentID = g.StudentID
	grade.SubjectID = g.SubjectID
	grade.Grade = g.Grade
	return grade, nil
}

func (s *GradeService) Update(id int, updatedGrade models.Grade) error {
	var g models.Grade
	db.DB.First(&g, id)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error fetching grade:", err)
		return err
	}
	g.Grade = updatedGrade.Grade
	g.StudentID = updatedGrade.StudentID
	g.SubjectID = updatedGrade.SubjectID
	db.DB.Save(&g)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error updating grade:", err)
		return err
	}
	return nil
}

func (s *GradeService) Delete(id int) error {
	var g models.Grade
	db.DB.Preload("Student").Preload("Subject").First(&g, id)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error fetching grade:", err)
		return err
	}

	db.DB.Unscoped().Delete(&g)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error deleting grade:", err)
		return err
	}
	return nil
}
