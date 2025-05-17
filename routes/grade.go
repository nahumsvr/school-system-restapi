package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nahumsvr/school-system-restapi/controllers"
	"github.com/nahumsvr/school-system-restapi/services"
)

func SetUpGradesRoutes(r *gin.Engine, s *services.GradeService) {
	grades := r.Group("/api/grades")
	gradeController := controllers.NewGradeController(s)

	grades.GET("/", gradeController.GetAll)
	grades.POST("/", gradeController.Create)
	grades.GET("/:grade_id", gradeController.Get)
	grades.PUT("/:grade_id", gradeController.Update)
	grades.DELETE("/:grade_id", gradeController.Delete)
}
