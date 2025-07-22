package database

import (
	"github.com/jmoiron/sqlx"
)

type Models struct {
	Users      UserModel
	Events     EventModel
	Attendence AttendenceModel
}

func NewModels(db *sqlx.DB) Models {
	return Models{
		Users:      UserModel{DB: db},
		Events:     EventModel{DB: db},
		Attendence: AttendenceModel{DB: db},
	}
}
