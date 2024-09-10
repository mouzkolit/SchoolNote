package models

import (
	"time"

	"github.com/uptrace/bun"
)

type SchoolLogin struct {
	ID         int64 `bun:",pk,autoincrement"`
	SchoolName string
	Password   string
}

type SchoolAdmin struct {
	ID       int64 `bun:",pk,autoincrement"`
	SchoolID int64 `bun:"type:int,notnull"`
	School   *School
	Name     string
	LastName string
	Email    string
	Password string
}

type School struct {
	ID        int64 `bun:",pk,autoincrement"`
	Name      string
	Location  string
	SchoolWeb string
	Classes   []*Class   `bun:"rel:has-many,join:id=school_id"`
	Teachers  []*Teacher `bun:"rel:has-many,join:id=school_id"`
	Pupils    []*Pupil   `bun:"rel:has-many,join:id=school_id"`
}

type Pupil struct {
	bun.BaseModel `bun:"table: pupil ,alias:pu"`
	ID            int64 `bun:",pk,autoincrement"`
	Name          string
	LastName      string
	Birthday      time.Time
	ClassID       int64 `bun:"type:int,notnull"` // Foreign key to Class
	SchoolID      int64 `bun:"type:int,notnull"` // Foreign key to School
}

type Teacher struct {
	bun.BaseModel `bun:"table: teacher ,alias:te"`
	ID            int64 `bun:",pk,autoincrement"`
	ClassID       int64 `bun:"type:int,notnull"` // Foreign key to Class
	SchoolID      int64 `bun:"type:int,notnull"` // Foreign key to School
	Name          string
	LastName      string
}

type Class struct {
	bun.BaseModel `bun:"table: class ,alias:cl"`
	ID            int64 `bun:",pk,autoincrement"`
	Class         string
	SchoolID      int64    `bun:"type:int,notnull"` // Foreign key to School
	TeacherID     int64    `bun:"type:int,notnull"` // Foreign key to Teacher
	Teacher       *Teacher `bun:"rel:belongs-to,join:teacher_id=id"`
	Pupils        []*Pupil `bun:"rel:has-many,join:id=class_id"`
}

type Course struct {
	ID      int64 `bun:",pk,autoincrement"`
	Name    string
	Teacher *Teacher
	Class   *Class
	Pupils  []*Pupil
}

type Note struct {
	ID      int64 `bun:",pk,autoincrement"`
	Teacher *Teacher
	Pupil   *Pupil
	Note    int
	Date    time.Time
}

type PupilDocumentation struct {
	ID            int64 `bun:",pk,autoincrement"`
	Documentation string
	PupilID       int64 `bun:"type:int,notnull"` // Foreign key to Pupil
}

type SchoolDocumentation struct {
	ID            int64 `bun:",pk,autoincrement"`
	Documentation string
	SchoolID      int64 `bun:"type:int,notnull"` // Foreign key to School
}

type Message struct {
	ID       int64 `bun:",pk,autoincrement"`
	Sender   string
	Receiver string
	Message  string
	Date     time.Time
}

type Document struct {
	ID      int64 `bun:",pk,autoincrement"`
	Title   string
	Content string
	Date    time.Time
}

type CourseAllocation struct {
	ID int64 `bun:",pk,autoincrement"`
}
