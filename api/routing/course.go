package routing

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mouzkolit/GOCli/database"
)

func CreateCourse(r *gin.Engine, db *database.DB) {
	r.POST("/school/:id/:class/course", func(c *gin.Context) {
		className := c.Param("class")
		schoolID := c.Param("id")
		courseName := c.Query("course")
		courseType := c.Query("Type")
		println("this is the course type: %s", courseType)
		println("this is the course name: %s", courseName)
		println("this is the class name %s", className)
		println("this is the school id %s", schoolID)
	})
}

func GetCourse(r *gin.Engine, db *database.DB) {
	r.GET("/school/:id/:class/course/:courseid", func(c *gin.Context) {
		courseID, err := strconv.ParseInt(c.Param("courseid"), 10, 64)
		if err != nil {
			println("error here %s", err)
		}
		className := c.Param("class")

		db.GetCourse(courseID, className)
	})
}
