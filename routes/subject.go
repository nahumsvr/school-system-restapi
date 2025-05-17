package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nahumsvr/school-system-restapi/controllers"
	"github.com/nahumsvr/school-system-restapi/services"
)

func SetUpSubjectsRoutes(r *gin.Engine, s *services.SubjectService) {
	subjects := r.Group("/api/subjects")
	subjectController := controllers.NewSubjectController(s)

	subjects.GET("/", subjectController.GetAll)
	subjects.POST("/", subjectController.Create)
	subjects.GET("/:subject_id", subjectController.Get)
	subjects.PUT("/", subjectController.Update)
	subjects.DELETE("/:subject_id", subjectController.Delete)
}
