package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	"github.com/mouzkolit/GOCli/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"golang.org/x/crypto/bcrypt"
)

type DB struct {
	db *bun.DB
}

// initializes a new database in memory
func NewDB() (*DB, error) {
	sqldb, err := sql.Open("sqlite3", "school.db")
	db := bun.NewDB(sqldb, sqlitedialect.New())
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) getContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

// Creates the base tables for the application
func (d *DB) InitSchema() {
	ctx := context.Background()
	models := []interface{}{
		(*models.Pupil)(nil),
		(*models.Teacher)(nil),
		(*models.Class)(nil),
		(*models.Course)(nil),
		(*models.Note)(nil),
		(*models.School)(nil),
		(*models.SchoolLogin)(nil),
	}

	for _, model := range models {
		exists, err := d.db.NewSelect().Model(model).Exists(ctx)
		if err != nil {
			log.Printf("Error checking if table exists for model %T: %v", model, err)
		}

		if !exists {
			_, err := d.db.NewCreateTable().Model(model).IfNotExists().Exec(ctx)
			if err != nil {
				log.Printf("Error creating table for model %T: %v", model, err)
			} else {
				log.Printf("Table created for model %T", model)
			}
		} else {
			log.Printf("Table already exists for model %T", model)
		}
	}
}

// Creates a new School in the shared database
func (d *DB) CreateSchool(name string, place string, web string) error {
	school := &models.School{
		Name:      name,
		Location:  place,
		SchoolWeb: web,
	}
	_, err := d.db.NewInsert().Model(school).Exec(d.getContext())
	if err != nil {
		return err
	}
	return nil
}

// Inserts a new Pupil in the shared database
func (d *DB) InsertPupil(name string, lastName string) {
	pupil := &models.Pupil{
		Name:     name,
		LastName: lastName,
		Birthday: time.Now(),
	}
	res, err := d.db.NewInsert().Model(pupil).Exec(d.getContext())
	if err != nil {
		panic(err)
	}
	println(res)

}

// Inserts a new Class in the shared database
func (d *DB) InsertClass(name string, schoolID int64) error {
	class := &models.Class{
		Class:    name,
		SchoolID: schoolID,
	}
	_, err := d.db.NewInsert().Model(class).Exec(d.getContext())
	if err != nil {
		return err
	}
	return nil
}

// Inserts a new Teacher in the shared database
func (d *DB) CreateTeacher(name string, lastName string, classID int64, schoolID int64) {
	teacher := &models.Teacher{
		Name:     name,
		LastName: lastName,
		ClassID:  classID,
		SchoolID: schoolID,
	}
	res, err := d.db.NewInsert().Model(teacher).Exec(d.getContext())
	if err != nil {
		panic(err)
	}
	println(res)
}

// Retrieves all Pupils from the database
func (d *DB) GetPupils() ([]models.Pupil, error) {
	var pupils []models.Pupil
	err := d.db.NewSelect().Model(&pupils).Scan(d.getContext())
	if err != nil {
		return nil, err
	}
	return pupils, nil
}

// Retrieves a Pupil from the database
func (d *DB) GetPupil(id int64) (models.Pupil, error) {
	var pupil models.Pupil
	err := d.db.NewSelect().Model(&pupil).Where("id = ?", id).Scan(d.getContext())
	if err != nil {
		return pupil, err
	}
	return pupil, nil
}

func (d *DB) AddLogin(model *models.SchoolLogin) error {
	_, err := d.db.NewInsert().Model(model).Exec(d.getContext())
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetSchool(id int64) (models.School, error) {
	var school models.School
	err := d.db.NewSelect().Model(&school).Where("id = ?", id).Scan(d.getContext())
	if err != nil {
		return school, err
	}
	return school, nil
}

func (d *DB) GetSchools() ([]models.School, error) {
	var schools []models.School
	err := d.db.NewSelect().Model(&schools).Scan(d.getContext())
	if err != nil {
		return nil, err
	}
	return schools, nil
}

func (d *DB) GetLogin(name string, password string) (bool, error) {
	var login models.SchoolLogin
	err := d.db.NewSelect().Model(&login).Where("school_name = ?", name).Scan(d.getContext())
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}
