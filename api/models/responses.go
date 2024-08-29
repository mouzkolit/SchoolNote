package models

import "time"

type PupilResponse struct {
	ID       int64
	Name     string
	LastName string
	Birthday time.Time
}

type SchoolResponse struct {
	ID        int64
	Name      string
	Location  string
	SchoolWeb string
}

type TeacherResponse struct {
	ID       int64
	Name     string
	LastName string
}

type ClassResponse struct {
	ID      int64
	Class   string
	Teacher *TeacherResponse
	Pupils  []*PupilResponse
}

type SchoolDocumentationResponse struct {
	ID            int64
	Documentation string
	SchoolID      int64
}
