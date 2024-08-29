package routing

import (
	"log"

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

		db.CreateSchool(name, schoolPlace, schoolWeb)
		err := CreateSchoolLogin(db, name)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error creating school login",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "pong",
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
