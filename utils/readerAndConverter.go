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

func ConvertToJson(c *gin.Context, body []byte, student *models.Student) {
	// err := json.NewEncoder(c.Writer).Encode(&student)
	err := json.Unmarshal(body, &student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parseando el JSON"})
		return
	}
}
