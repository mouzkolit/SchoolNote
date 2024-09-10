package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/mouzkolit/GOCli/database"
)

func AddCourse(r *gin.Engine, db *database.DB) {
	r.POST("/:class/course", func(c *gin.Context) {
		className := c.Param("class")
		courseName := c.Query("course")
		courseType := c.Query("Type")
		println("this is the course type: %s", courseType)
		println("this is the course name: %s", courseName)
		println("this is the class name %s", className)
	})
}

func GetCourse(r *gin.Engine, db *database.DB) {
	r.GET("/:class/course/:courseid", func(c *gin.Context) {
		courseID := c.Param("courseid")
		className := c.Param("class")
	})
}
