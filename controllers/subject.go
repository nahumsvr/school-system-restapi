package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nahumsvr/school-system-restapi/models"
	"github.com/nahumsvr/school-system-restapi/services"
	"github.com/nahumsvr/school-system-restapi/utils"
)

type SubjectController struct {
	SubjectService *services.SubjectService
}

func NewSubjectController(subjectService *services.SubjectService) *SubjectController {
	return &SubjectController{
		SubjectService: subjectService,
	}
}

func (s *SubjectController) GetAll(c *gin.Context) {
	users := s.SubjectService.GetAll()
	c.JSON(200, users)
}

func (s *SubjectController) Create(c *gin.Context) {
	body := utils.ReadBody(c)
	var subject models.Subject
	utils.SubjectConvertToJson(c, body, &subject)

	if subject.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El nombre son obligatorios"})
		return
	}

	newSubject := s.SubjectService.Create(subject)
	c.JSON(http.StatusOK, newSubject)
}

func (s *SubjectController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("subject_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	subject, subjectErr := s.SubjectService.Get(id)
	if subjectErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": subjectErr.Error()})
		return
	}
	c.JSON(http.StatusOK, subject)
}

func (s *SubjectController) Update(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("subject_id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
	// 	return
	// }
	var subject models.Subject
	utils.SubjectConvertToJson(c, utils.ReadBody(c), &subject)
	newSubject, subjectErr := s.SubjectService.Update(int(subject.SubjectID), subject)
	if subjectErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": subjectErr.Error()})
		return
	}
	c.JSON(http.StatusOK, newSubject)
}

func (s *SubjectController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("subject_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	subjectErr := s.SubjectService.Delete(id)
	if subjectErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": subjectErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
