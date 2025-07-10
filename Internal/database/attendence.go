package database

import "database/sql"

type AttendenceModel struct {
	DB *sql.DB
}

type Attendence struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}
