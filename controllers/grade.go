package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nahumsvr/school-system-restapi/models"
	"github.com/nahumsvr/school-system-restapi/services"
	"github.com/nahumsvr/school-system-restapi/utils"
)

type GradeController struct {
	GradeService *services.GradeService
}

func NewGradeController(gradeService *services.GradeService) *GradeController {
	return &GradeController{
		GradeService: gradeService,
	}
}

func (g *GradeController) GetAll(c *gin.Context) {
	grades := g.GradeService.GetAll()
	c.JSON(200, grades)
}

func (g *GradeController) Create(c *gin.Context) {
	body := utils.ReadBody(c)
	var grade models.Grade
	utils.GradeConvertToJson(c, body, &grade)

	if grade.Grade < 0 || grade.Grade > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La calificación debe estar entre 0 y 100"})
		return
	}

	newGrade := g.GradeService.Create(grade)
	c.JSON(http.StatusOK, newGrade)
}

func (g *GradeController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	grade, gradeErr := g.GradeService.Get(id)
	if gradeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": gradeErr.Error()})
		return
	}
	c.JSON(http.StatusOK, grade)
}

func (g *GradeController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	body := utils.ReadBody(c)
	var grade models.Grade
	utils.GradeConvertToJson(c, body, &grade)
	updatedGrade, updateErr := g.GradeService.Update(id, grade)
	if updateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": updateErr.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedGrade)
}

func (g *GradeController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número"})
		return
	}
	deleteErr := g.GradeService.Delete(id)
	if deleteErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": deleteErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Grade deleted"})
}
