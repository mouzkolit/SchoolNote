package routing

import (
	"log"
	"net/http"
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
		pwd, err := CreateSchoolLogin(db, name)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error creating school login",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "School Successfully created with password: " + pwd,
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

func SchoolLogin(r *gin.Engine, db *database.DB) {
	r.POST("/school/login", func(c *gin.Context) {
		name := c.Query("name")
		password := c.Query("password")
		success, err := CheckSchoolLogin(db, name, password)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error checking school login",
			})
			return
		}
		if success {
			// Generate an access token (you'll need to implement this)
			accessToken, err := generateAccessToken(name)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Error generating access token",
				})
				return
			}

			// Set the cookie
			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie(
				"access_token",
				accessToken,
				3600,  // Max age in seconds (1 hour)
				"/",   // Path
				"",    // Domain
				false, // Secure (set to true if using HTTPS)
				true,  // HttpOnly
			)

			c.JSON(200, gin.H{
				"message": "Login successful",
			})
		} else {
			c.JSON(401, gin.H{
				"message": "Invalid credentials",
			})
		}
	})
}

// Add this function to generate an access token
func generateAccessToken(schoolName string) (string, error) {
	// Implement your token generation logic here
	// This is just a placeholder
	return "sample_access_token_for_" + schoolName, nil
}

func CheckSchoolLogin(db *database.DB, name string, password string) (bool, error) {
	success, err := db.GetLogin(name, password)
	if err != nil {
		log.Printf("Login check failed: %v", err)
		return success, err
	}
	return success, nil
}

func GenerateAccessToken() (string, error) {
	// we need to implement here logic around the oAuth scheme
	return "herewego", nil
}

func CreateSchoolLogin(db *database.DB, name string) (string, error) {
	autoPwd, err := generateRandomPassword()
	if err != nil {
		// here needs to be a return statement
		return "", err
	}
	// Create a new school login
	new_model := &models.SchoolLogin{
		SchoolName: name,
		Password:   autoPwd,
	}
	db.AddLogin(new_model)
	return autoPwd, nil
}

func generateRandomPassword() (string, error) {
	res, err := password.Generate(20, 10, 10, false, false)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return res, nil
}
