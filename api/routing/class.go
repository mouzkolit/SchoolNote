package routing

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mouzkolit/GOCli/database"
)

// PostPupil             godoc
// @Summary      Create School Class
// @Description  Creates a new school class in the system
// @Tags         pupil
// @Produce      json
// @Param        schoolID path string true "LastName of the Pupil"
// @Param        name path string true "Name of the Pupil"
// @Success      200   {object}  models.PupilResponse
// @Router       /pupil [post]
func CreateClass(r *gin.RouterGroup, db *database.DB) {
	r.POST("/:school_id/:class", func(c *gin.Context) {
		name := c.Param("name")
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
			"message": "Successfully added Class to School",
		})
	})
}
