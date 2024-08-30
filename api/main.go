package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	db, err := InitializeDB()
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = false
	config.AllowOrigins = append(config.AllowOrigins, "http://localhost:5173")
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routing.CreatePupil(r, db)
	routing.GetPupil(r, db)
	routing.GetPupils(r, db)
	routing.CreateSchool(r, db)
	routing.GetSchool(r, db)
	routing.GetSchools(r, db)
	routing.SchoolLogin(r, db)
	r.Run(":8080")
}

// func setupCors() gin.HandlerFunc {
// 	config := cors.DefaultConfig()
// 	config.AllowOrigins = []string{"http://localhost:5173"} // Your frontend URL
// 	config.AllowCredentials = true
// 	return cors.New(config)
// }

// InitializeDB initializes the Sqlite database for now
// this will be replaced with a postgres database sitting on Azure Cloud
func InitializeDB() (*database.DB, error) {
	db, err := database.NewDB()
	if err != nil {
		fmt.Println(err)
	}
	db.InitSchema()
	return db, err
}
