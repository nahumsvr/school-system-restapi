package services

import (
	"fmt"

	"github.com/nahumsvr/school-system-restapi/db"
	"github.com/nahumsvr/school-system-restapi/models"
)

type SubjectService struct {
	subject []models.Subject
}

func NewSubjectService() *SubjectService {
	return &SubjectService{
		subject: []models.Subject{},
	}
}

func (s *SubjectService) GetAll() []models.Subject {
	subject := db.DB.Find(&s.subject)
	err := subject.Error
	if err != nil {
		fmt.Println("Error fetching subject:", err)
		return nil
	}
	return s.subject
}

func (s *SubjectService) Create(subject models.Subject) models.Subject {
	fmt.Println("Creating subject:", subject)
	createdSubject := db.DB.Create(&subject)
	err := createdSubject.Error
	if err != nil {
		fmt.Println("Error creating subject:", err)
		return models.Subject{}
	}
	return subject
}

func (s *SubjectService) Get(id int) (subject models.Subject, err error) {
	db.DB.First(&subject, id)
	if err := db.DB.Error; err != nil {
		fmt.Println("Error fetching subject:", err)
		return models.Subject{}, err
	}
	return subject, nil
}

func (s *SubjectService) Update(id int, updatedSubject models.Subject) (subject models.Subject, err error) {
	if err := db.DB.First(&subject, id).Error; err != nil {
		fmt.Println("Error finding subject:", err)
		return models.Subject{}, err
	}

	if err := db.DB.Model(&subject).Updates(updatedSubject).Error; err != nil {
		fmt.Println("Error updating subject:", err)
		return models.Subject{}, err
	}

	return subject, nil
}

func (s *SubjectService) Delete(id int) error {
	subject, err := s.Get(id)
	if err != nil {
		fmt.Println("Error fetching subject:", err)
		return err
	}
	db.DB.Unscoped().Delete(&subject)
	return nil
}
