package routing

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mouzkolit/GOCli/database"
)

// PostPupil             godoc
// @Summary      Creates A Pupil in a Class
// @Description  Creates A Pupil in a Class
// @Tags         pupil
// @Produce      json
// @Param        name query string true "Name of the Pupil"
// @Param        lastName query string true "LastName of the Pupil"
// @Success      200   {object}  models.PupilResponse
// @Router       /pupil [post]
func CreatePupil(r *gin.Engine, db *database.DB) {
	r.POST("/pupil", func(c *gin.Context) {
		name := c.Query("name")
		lastName := c.Query("age")

		db.InsertPupil(name, lastName)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

// UploadPupil godoc
func CreatePupilFromExcel(r *gin.Engine, db *database.DB) {
	r.POST("/pupil/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fileExt := strings.ToLower(filepath.Ext(file.Filename))
		println(fileExt)
	})
}

// GetPupil             godoc
// @Summary      Retrieves a pupil from the path
// @Description  Gets a Pupil from a Class
// @Tags         Pupil
// @Param        id path int true "ID of the Pupil"
// @Produce      json
// @Success      200   {object}  models.PupilResponse
// @Router       /pupil/{id} [get]
func GetPupil(r *gin.Engine, db *database.DB) {
	r.GET("/pupil/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Wrong ID format",
			})
			return
		}

		resp, err := db.GetPupil(id)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "No User found",
			})
			return
		}

		c.JSON(200, gin.H{
			"data": resp,
		})
	})
}

// GetPupils             godoc
// @Summary      Retrieves all Pupils
// @Description  Retrieves all Pupils from the school
// @Tags         Pupils
// @Produce      json
// @Success      200   {object}  []models.PupilResponse
// @Router       /pupils [get]
func GetPupils(r *gin.Engine, db *database.DB) {
	r.GET("/pupils", func(c *gin.Context) {
		jsonData, err := db.GetPupils()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "No User found",
			})
		}
		c.JSON(200, gin.H{"data": jsonData})
	})
}
