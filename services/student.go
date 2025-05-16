package services

import (
	"fmt"

	"github.com/nahumsvr/school-system-restapi/db"
	"github.com/nahumsvr/school-system-restapi/models"
)

type StudentService struct {
	students []models.Student
}

func NewStudentService() *StudentService {
	return &StudentService{
		students: []models.Student{},
	}
}

func (s *StudentService) GetAll() []models.Student {
	students := db.DB.Find(&s.students)
	err := students.Error
	if err != nil {
		fmt.Println("Error fetching students:", err)
		return nil
	}
	return s.students
}

func (s *StudentService) Create(student models.Student) models.Student {
	fmt.Println("Creating student:", student)
	createdStudent := db.DB.Create(&student)
	err := createdStudent.Error
	if err != nil {
		fmt.Println("Error creating student:", err)
		return models.Student{}
	}
	return student
}

func (s *StudentService) Get(id int) (student models.Student, err error) {
	db.DB.First(&student, id)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error fetching student:", err)
		return models.Student{}, err
	}
	return student, nil
}

func (s *StudentService) Update(id int, updatedStudent models.Student) (student models.Student, err error) {
	if err := db.DB.First(&student, id).Error; err != nil {
		fmt.Println("Error finding student:", err)
		return models.Student{}, err
	}

	if err := db.DB.Model(&student).Updates(updatedStudent).Error; err != nil {
		fmt.Println("Error updating student:", err)
		return models.Student{}, err
	}

	return student, nil
}

func (s *StudentService) Delete(id int) error {
	student, err := s.Get(id)
	if err != nil {
		fmt.Println("Error fetching student:", err)
		return err
	}
	db.DB.Unscoped().Delete(&student)
	return nil
}
