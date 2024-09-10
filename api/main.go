package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mouzkolit/GOCli/database"
	_ "github.com/mouzkolit/GOCli/docs"
	"github.com/mouzkolit/GOCli/routing"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/// this should be an applciation for teachers to take notes and save notes
// for students to view, can also be shared with parents

// @title           School Notebook API
// @version         0.1.0
// @description     A School Notes service API Backend in GOlang

// @contact.name   Maximilian Zeidler PhD
// @contact.email  maximilian.zeidler@outlook.com

// @host      localhost:8080
// @BasePath  /
func main() {
	db, err := database.InitializeDB()
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	config, err := initConfig()
	r.Use(cors.New(config))
	public := r.Group("/")
	private := r.Group("/")
	private.Use(checkJWT())
	// enable the swagger ui for testing
	public.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//public endpoints
	routing.CreateSchool(public, db)
	routing.SchoolLogin(public, db)
	// routing will be performmed here

	//private endpoints
	routing.CreatePupil(r, db)
	routing.GetPupil(r, db)
	routing.GetPupils(r, db)

	routing.GetSchool(r, db)
	routing.GetSchools(private, db)

	routing.CreateClass(private, db)

	routing.GetCourse(r, db)
	routing.CreateCourse(r, db)

	r.Run(":8080")
}

func initConfig() (cors.Config, error) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:8080"} // Remove the "*" at the end
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
	config.ExposeHeaders = []string{"Content-Length", "Set-Cookie"}
	config.AllowCredentials = true
	return config, nil
}

func checkJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("access_token")
		if tokenString == "" || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// here the token should be validated
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("hellohereweare"), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
