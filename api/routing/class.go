package routing

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mouzkolit/GOCli/database"
)

// PostClass             godoc
// @Summary      Create School Class
// @Description  Creates a new school class in the system
// @Tags         pupil
// @Produce      json
// @Param        schoolID path string true "Schoool ID"
// @Param        class path string for the class to crate one
// @Success      200   {object}  models.PupilResponse
// @Router       /pupil [post]
func CreateClass(r *gin.RouterGroup, db *database.DB) {
	r.POST("/:school_id/:class", func(c *gin.Context) {
		name := c.Param("class")
		schoolID, err := strconv.ParseInt(c.Param("schoolID"), 10, 64)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Wrong ID format",
			})
			return
		}

		err = db.InsertClass(name, schoolID)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Error inserting class",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": "Successfully added Class to School",
		})
	})
}

// GetClass             godoc
// @Summary      Retrieves all school classes
// @Description  Retrieves all school classes for the school
// @Tags         class
// @Produce      json
// @Param        schoolID path string true "LastName of the Pupil"
// @Param        name path string true "Name of the Pupil"
// @Success      200   {object}  models.PupilResponse
// @Router       /pupil [post]
func GetClasses(r *gin.RouterGroup, db *database.DB) {
	r.POST("/:school_id/classes", func(c *gin.Context) {
		schoolID, err := strconv.ParseInt(c.Param("schoolID"), 10, 64)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Wrong ID format",
			})
			return
		}

		classes, err := db.GetClasses(schoolID)
		if err != nil {
			c.JSON(404, gin.H{
				"message": fmt.Sprintf("Error Fetching classes with error %s", err),
			})
			return
		}
		c.JSON(200, gin.H{
			"data": classes,
		})
	})
}
