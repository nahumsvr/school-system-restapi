package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nahumsvr/school-system-restapi/db"
	"github.com/nahumsvr/school-system-restapi/models"
	"github.com/nahumsvr/school-system-restapi/routes"
	"github.com/nahumsvr/school-system-restapi/services"
)

func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.Student{})
	db.DB.AutoMigrate(&models.Subject{})
	db.DB.AutoMigrate(&models.Grade{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.SetUpStudentRoutes(r, &services.StudentService{})
	routes.SetUpSubjectsRoutes(r, &services.SubjectService{})
	routes.SetUpGradesRoutes(r, &services.GradeService{})
	r.Run(":3000")

	fmt.Println("Servidor escuchando en el puerto 3000")
	fmt.Println("http://localhost:3000")
}
