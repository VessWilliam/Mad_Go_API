package database

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type AttendenceModel struct {
	DB *sql.DB
}

type Attendence struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}

func (m *AttendenceModel) Insert(attendence *Attendence) (*Attendence, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO attendences (event_id, user_id) VALUES ($1, $2) RETURNING id"
	log.Printf("Inserting attendence: event_id=%d, user_id=%d", attendence.EventId, attendence.UserId)

	err := m.DB.QueryRowContext(ctx, query, attendence.EventId, attendence.UserId).Scan(&attendence.Id)
	if err != nil {
		log.Printf("Insert Attendence Error: %v", err)
		return nil, err
	}

	log.Printf("Attendence inserted with ID: %d", attendence.Id)
	return attendence, nil
}

func (m *AttendenceModel) GetByEventAndAttendence(eventId, userId int) (*Attendence, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM attendences WHERE event_id = $1 AND user_id = $2"
	log.Printf("Querying attendence with event_id=%d and user_id=%d", eventId, userId)

	var attendence Attendence
	err := m.DB.QueryRowContext(ctx, query, eventId, userId).Scan(
		&attendence.Id,
		&attendence.UserId,
		&attendence.EventId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No attendence found for event_id=%d and user_id=%d", eventId, userId)
			return nil, nil
		}
		log.Printf("Error querying attendence: %v", err)
		return nil, err
	}

	log.Printf("Found attendence: %+v", attendence)
	return &attendence, nil
}

func (m *AttendenceModel) GetAttendenceByEvent(eventId int) ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT u.id, u.name, u.email
		FROM users u
		JOIN attendences a ON u.id = a.user_id
		WHERE a.event_id = $1
	`

	rows, err := m.DB.QueryContext(ctx, query, eventId)
	if err != nil {
		log.Printf("Query error in GetAttendenceByEvent: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			log.Printf("Row scan error in GetAttendenceByEvent: %v", err)
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error in GetAttendenceByEvent: %v", err)
		return nil, err
	}

	return users, nil
}
