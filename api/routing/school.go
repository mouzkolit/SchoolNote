package routing

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mouzkolit/GOCli/database"
	"github.com/mouzkolit/GOCli/models"
	"golang.org/x/crypto/bcrypt"
)

type Bearer struct {
	AccessToken string `json:"access_token"`
}

type SchoolInfo struct {
	Username   string
	Role       string
	SchoolName string
	SchoolID   int
}

// CreateSchool             godoc
// @Summary      Create School
// @Description  Creates a new school in the system
// @Tags         School
// @Produce      json
// @Param        name query string true "Name of the School"
// @Param        place query string true "Place of the School"
// @Param        web query string true "Website of the School"
// @Success      200   {object}  models.SchoolResponse
// @Router       /school [post]
func CreateSchool(r *gin.RouterGroup, db *database.DB) {
	r.POST("/school", func(c *gin.Context) {
		name := c.Query("name")
		schoolPlace := c.Query("place")
		schoolWeb := c.Query("web")
		password := c.Query("password")

		err := db.CreateSchool(name, schoolPlace, schoolWeb)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error creating school",
			})
			return
		}
		pwd, err := CreateSchoolLogin(db, name, password)
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

// GetSchool             godoc
// @Summary      Get School
// @Description  Gets a school by ID
// @Tags         School
// @Produce      json
// @Param        id path string true "ID of the School"
// @Success      200   {object}  models.SchoolResponse
// @Router       /school/:id [get]
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

// GetSchools             godoc
// @Summary      Create School Class
// @Description  Creates a new school class in the system
// @Tags         School
// @Produce      json
// @Success      200   {object}  models.SchoolResponse
// @Router       /schools [get]
func GetSchools(r *gin.RouterGroup, db *database.DB) {
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

// SchoolLogin             godoc
// @Summary      School Login
// @Description  Logs in a school
// @Tags         School
// @Param        name query string true "Name of the School"
// @Param        password query string true "Password of the School"
// @Produce      json
// @Success      200   {object}  models.SchoolResponse
// @Router       /school/login [post]
func SchoolLogin(r *gin.RouterGroup, db *database.DB) {
	r.POST("/school/login", func(c *gin.Context) {
		name := c.Query("name")
		password := c.Query("password")
		success := CheckSchoolLogin(db, name, password)
		if success.err != nil {
			c.JSON(500, gin.H{
				"message": "Error checking school login",
			})
			return
		}

		// Set the cookie
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie(
			"access_token",
			success.token,
			3600,        // Max age in seconds (1 hour)
			"/",         // Path
			"localhost", // Domain
			false,       // Secure (set to true if using HTTPS)
			true,        // HttpOnly
		)

		c.JSON(200, gin.H{
			"message": "Login successful",
		})
	})
}

type loginErr struct {
	token string
	err   error
}

// Add this function to generate an access token
func CheckSchoolLogin(db *database.DB, name string, password string) loginErr {
	success, err := db.GetLogin(name, password)
	if err != nil {
		log.Printf("Login check failed: %v", err)
		return loginErr{
			token: "",
			err:   err,
		}
	}

	if success {
		token, err := generateAccessToken(name)
		if err != nil {
			log.Printf("Login check failed: %v", err)
			return loginErr{
				token: "",
				err:   err,
			}
		}
		return loginErr{
			token: token,
			err:   nil,
		}
	}

	return loginErr{
		token: "",
		err:   errors.New("Invalid Credentials"),
	}
}

func generateAccessToken(schoolName string) (string, error) {
	// we need to implement here logic around the oAuth scheme
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"schoolName": schoolName,
			"exp":        time.Now().Add(time.Hour * 2).Unix(),
		})

	secretKey := []byte("hellohereweare")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateSchoolLogin(db *database.DB, name string, password string) (string, error) {
	autoPwd, err := generateHashFromPassword(password)
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

func generateHashFromPassword(password string) (string, error) {
	// hash the password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(hashedPwd), nil
}
