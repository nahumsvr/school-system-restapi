package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/nahumsvr/school-system-restapi/controller"
	"github.com/nahumsvr/school-system-restapi/services"
)

func SetUpStudentRoutes(r *gin.Engine, s *services.StudentService) {
	students := r.Group("/api/students")
	studentController := controllers.NewStudentController(s)

	students.GET("/", studentController.GetAll)
	students.POST("/", studentController.Create)
	students.GET("/:student_id", studentController.Get)
	students.PUT("/:student_id", studentController.Update)
	students.DELETE("/:student_id", studentController.Delete)
}
