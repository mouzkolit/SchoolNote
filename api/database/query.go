package database

import (
	"fmt"
	"time"

	"github.com/mouzkolit/GOCli/models"
	"golang.org/x/crypto/bcrypt"
)

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

func (d *DB) GetClasses(schoolID int64) ([]models.Class, error) {
	var classes []models.Class
	err := d.db.NewSelect().Model(&classes).Where("school_id = ?", schoolID)
	if err != nil {
		return classes, fmt.Errorf("iderror: no classes found for this id")
	}

	return classes, nil
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

func (d *DB) GetCourse(id int64, className string) (*models.Course, error) {
	var course models.Course
	err := d.db.NewSelect().Model(&course).Where("id = ?", id)
	println(err)
	return &course, nil
}
