package models

import (
	"errors"
	"time"

	"dkds.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO event
	(name, description, location, dateTime, userId)
	VALUES
	(?, ?, ?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("Could not save event, " + err.Error())
	}

	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return errors.New("Could not save event, " + err.Error())
	}
	defer statement.Close()

	id, err := result.LastInsertId()
	if err != nil {
		return errors.New("Could not retrieve last saved ID, " + err.Error())
	}

	e.ID = id

	return nil
}

func (e *Event) Update(id int64) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return errors.New("Could not update event, " + err.Error())
	}

	_, err = GetEventById(id)
	if err != nil {
		return errors.New("Could not update event, " + err.Error())
	}

	query := `
	UPDATE event
	SET
		name=?, 
		description=?, 
		location=?, 
		dateTime=?, 
		userId=?
	WHERE 
		id = ?
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("Could not update event, " + err.Error())
	}

	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, id)
	if err != nil {
		return errors.New("Could not update event, " + err.Error())
	}
	defer statement.Close()

	err = tx.Commit()
	if err != nil {
		return errors.New("Could not update event, " + err.Error())
	}
	e.ID = id

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT id, name, description, location, dateTime, userId
	FROM event`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New("Could not retrieve events, " + err.Error())
	}
	defer rows.Close()

	var events = []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		)
		if err != nil {
			return nil, errors.New("Could not read event, " + err.Error())
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
		SELECT id, name, description, location, dateTime, userId
		FROM event
		WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID,
	)
	if err != nil {
		return nil, errors.New("Could not retrieve the event" + err.Error())
	}

	return &event, nil
}
