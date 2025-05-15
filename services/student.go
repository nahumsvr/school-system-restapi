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
	for _, student := range s.students {
		if int(student.StudentID) == id {
			return student, nil
		}
	}
	return models.Student{}, fmt.Errorf("student with ID %d not found", id)
}

func (s *StudentService) Update(id int, student models.Student) (models.Student, error) {
	for i, student := range s.students {
		if int(student.StudentID) == id {
			s.students[i].Name = student.Name
			s.students[i].Group = student.Group
			s.students[i].Email = student.Email
			return s.students[i], nil
		}
	}
	return models.Student{}, fmt.Errorf("student with ID %d not found", id)
}

func (s *StudentService) Delete(id int) error {
	for i, student := range s.students {
		if int(student.StudentID) == id {
			s.students = append(s.students[:i], s.students[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("student with ID %d not found", id)
}
