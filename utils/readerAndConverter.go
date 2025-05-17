package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahumsvr/school-system-restapi/models"
)

func ReadBody(c *gin.Context) []byte {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el body"})
		return []byte{0}
	}
	return body
}

func StudentConvertToJson(c *gin.Context, body []byte, student *models.Student) {
	err := json.Unmarshal(body, &student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parseando el JSON"})
		return
	}
}

func SubjectConvertToJson(c *gin.Context, body []byte, subject *models.Subject) {
	err := json.Unmarshal(body, &subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parseando el JSON"})
		return
	}
}

func GradeConvertToJson(c *gin.Context, body []byte, grade *models.Grade) {
	err := json.Unmarshal(body, &grade)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parseando el JSON"})
		return
	}
}
