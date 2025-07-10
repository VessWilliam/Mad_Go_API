package database

import "database/sql"

type Models struct {
	Users      UserModel
	Events     EventModel
	Attendence AttendenceModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:      UserModel{DB: db},
		Events:     EventModel{DB: db},
		Attendence: AttendenceModel{DB: db},
	}
}
