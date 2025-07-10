package database

import (
	"database/sql"
	utils "rest_api_gin/Internal/Utils"
)

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	Id          int             `json:"id"`
	OwnerId     int             `json:"ownerId" binding:"required"`
	Name        string          `json:"name" binding:"required,min=3"`
	Description string          `json:"desc" binding:"required,min=10"`
	Date        utils.DateSlash `json:"date" binding:"required"`
	Location    string          `json:"location" binding:"required, min=3"`
}
