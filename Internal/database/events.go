package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	utils "rest_api_gin/internal/Utils"
	"time"
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
	Location    string          `json:"location" binding:"required,min=3"`
}

func (m *EventModel) Insert(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println(event)

	query := "INSERT INTO events (owner_id, name, description, date, location) VALUES (?, ?, ?, ?, ?) RETURNING id"

	err := m.DB.QueryRowContext(ctx, query,
		event.OwnerId,
		event.Name,
		event.Description,
		event.Date.Time,
		event.Location,
	).Scan(&event.Id)
	if err != nil {
		log.Println("Insert Event Error:", err)
		return err
	}

	return nil
}

func (m *EventModel) GetAll() ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		SELECT id,
		 owner_id, 
		 name, 
		 description,
		 date, 
		 location
		FROM events
		ORDER BY date ASC
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Get Event list Error:", err)
		return nil, err
	}

	defer rows.Close()

	events := []*Event{}

	for rows.Next() {
		var event Event

		if err := rows.Scan(&event.Id,
			&event.OwnerId,
			&event.Name,
			&event.Description,
			&event.Date.Time,
			&event.Location); err != nil {
			log.Println("GetAll: Scan error:", err)
			return nil, err
		}

		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		log.Println("GetAll: Rows iteration error:", err)
		return nil, err
	}

	return events, nil
}

func (m *EventModel) Get(id int) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select  
    	id, 
		owner_id, 
		name, 
		description, 
		date, 
		location
	 from events where id = ?`

	var event Event

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&event.Id,
		&event.OwnerId,
		&event.Name,
		&event.Description,
		&event.Date.Time,
		&event.Location,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Get By ID Error:", err)
		return nil, err
	}

	return &event, nil
}

func (m *EventModel) Update(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "Update events Set name = ?, description = ? , date = ?, location = ? where id = ?"

	_, err := m.DB.ExecContext(ctx, query, event.Name, event.Description, event.Date.Time, event.Location, event.Id)
	if err != nil {
		log.Printf("Update Event failed for ID %d: %v\n", event.Id, err)
		return err
	}
	return nil
}

func (m *EventModel) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "Delete from events where id = ?"

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Delete Event Error (ID %d): %v\n", id, err)
		return err
	}
	return nil
}
