package routing

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mouzkolit/GOCli/database"
	"github.com/mouzkolit/GOCli/models"
	"github.com/sethvargo/go-password/password"
)

func CreateSchool(r *gin.Engine, db *database.DB) {
	r.POST("/school", func(c *gin.Context) {
		name := c.Query("name")
		schoolPlace := c.Query("place")
		schoolWeb := c.Query("web")

		err := db.CreateSchool(name, schoolPlace, schoolWeb)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error creating school",
			})
			return
		}
		err = CreateSchoolLogin(db, name)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error creating school login",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "School Successfully created",
		})
	})
}

func GetSchool(r *gin.Engine, db *database.DB) {
	r.GET("/school/:id", func(c *gin.Context) {
		schoolId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Wrong ID format",
			})
			return
		}
		school, err := db.GetSchool(schoolId)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error getting school",
			})
			return
		}
		c.JSON(200, gin.H{
			"school": school,
		})
	})
}

func GetSchools(r *gin.Engine, db *database.DB) {
	r.GET("/schools", func(c *gin.Context) {
		schools, err := db.GetSchools()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error getting schools",
			})
			return
		}
		c.JSON(200, gin.H{
			"schools": schools,
		})
	})
}

func CreateSchoolLogin(db *database.DB, name string) error {
	autoPwd, err := generateRandomPassword()
	if err != nil {
		// here needs to be a return statement
		return err
	}
	// Create a new school login
	new_model := &models.SchoolLogin{
		SchoolName: name,
		Password:   autoPwd,
	}
	db.AddLogin(new_model)
	return nil
}

func generateRandomPassword() (string, error) {
	// Generate a random password
	// Implement your logic here to generate a random password
	// For example, you can use a library like "github.com/sethvargo/go-password/password" to generate a random password
	// Make sure to import the necessary package and call the appropriate function
	res, err := password.Generate(20, 10, 10, false, false)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return res, nil
}
