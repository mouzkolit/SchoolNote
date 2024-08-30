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
)

type DB struct {
	db *bun.DB
}

// initializes a new database in memory
func NewDB() (*DB, error) {
	sqldb, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	db := bun.NewDB(sqldb, sqlitedialect.New())
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

func (d *DB) Close() error {
	return d.db.Close()
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
		res, err := d.db.NewCreateTable().Model(model).Exec(ctx)
		if err != nil {
			log.Fatalf("Failed to create table for model %T: %v", model, err)
		}
		log.Printf("Table created for model %T: %v", model, res)
	}

}

// Creates a new School in the shared database
func (d *DB) CreateSchool(name string, place string, web string) error {
	ctx := context.Background()
	school := &models.School{
		Name:      name,
		Location:  place,
		SchoolWeb: web,
	}
	_, err := d.db.NewInsert().Model(school).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Inserts a new Pupil in the shared database
func (d *DB) InsertPupil(name string, lastName string) {
	ctx := context.Background()
	pupil := &models.Pupil{
		Name:     name,
		LastName: lastName,
		Birthday: time.Now(),
	}
	res, err := d.db.NewInsert().Model(pupil).Exec(ctx)
	if err != nil {
		panic(err)
	}
	println(res)

}

// Inserts a new Class in the shared database
func (d *DB) InsertClass(name string, schoolID int64) error {
	ctx := context.Background()
	class := &models.Class{
		Class:    name,
		SchoolID: schoolID,
	}
	_, err := d.db.NewInsert().Model(class).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Inserts a new Teacher in the shared database
func (d *DB) CreateTeacher(name string, lastName string, classID int64, schoolID int64) {
	ctx := context.Background()
	teacher := &models.Teacher{
		Name:     name,
		LastName: lastName,
		ClassID:  classID,
		SchoolID: schoolID,
	}
	res, err := d.db.NewInsert().Model(teacher).Exec(ctx)
	if err != nil {
		panic(err)
	}
	println(res)
}

// Retrieves all Pupils from the database
func (d *DB) GetPupils() ([]models.Pupil, error) {
	ctx := context.Background()
	var pupils []models.Pupil
	err := d.db.NewSelect().Model(&pupils).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return pupils, nil
}

// Retrieves a Pupil from the database
func (d *DB) GetPupil(id int64) (models.Pupil, error) {
	ctx := context.Background()
	var pupil models.Pupil
	err := d.db.NewSelect().Model(&pupil).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return pupil, err
	}
	return pupil, nil
}

func (d *DB) AddLogin(model *models.SchoolLogin) error {
	ctx := context.Background()
	_, err := d.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetSchool(id int64) (models.School, error) {
	ctx := context.Background()
	var school models.School
	err := d.db.NewSelect().Model(&school).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return school, err
	}
	return school, nil
}

func (d *DB) GetSchools() ([]models.School, error) {
	ctx := context.Background()
	var schools []models.School
	err := d.db.NewSelect().Model(&schools).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return schools, nil
}

func (d *DB) GetLogin(name string, password string) (bool, error) {
	ctx := context.Background()
	var login models.SchoolLogin
	err := d.db.NewSelect().Model(&login).Where("name = ?", name).Where("password = ?", password).Scan(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
