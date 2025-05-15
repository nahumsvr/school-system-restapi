package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nahumsvr/school-system-restapi/models"
	"github.com/nahumsvr/school-system-restapi/services"
	"github.com/nahumsvr/school-system-restapi/utils"
)

type StudentController struct {
	StudentService *services.StudentService
}

func NewStudentController(studentService *services.StudentService) *StudentController {
	return &StudentController{
		StudentService: studentService,
	}
}

func (s *StudentController) GetAll(c *gin.Context) {
	users := s.StudentService.GetAll()
	c.JSON(200, users)
}

func (s *StudentController) Create(c *gin.Context) {
	body := utils.ReadBody(c)
	var student models.Student
	utils.ConvertToJson(c, body, &student)

	if student.Name == "" || student.Group == "" || student.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El nombre, grupo y correo son obligatorios"})
		return
	}

	newStudent := s.StudentService.Create(student)
	c.JSON(http.StatusOK, newStudent)
}

func (s *StudentController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	student, studentErr := s.StudentService.Get(id)
	if studentErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": studentErr.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (s *StudentController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	body := utils.ReadBody(c)
	var student models.Student
	utils.ConvertToJson(c, body, &student)
	newStudent, studentErr := s.StudentService.Update(id, student)
	if studentErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": studentErr.Error()})
		return
	}
	c.JSON(http.StatusOK, newStudent)
}

func (s *StudentController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	studentErr := s.StudentService.Delete(id)
	if studentErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": studentErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
