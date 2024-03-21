package main

import "time"

type MessageRequest struct {
	CreatedAt time.Time `json:"createdAt"`
	TEXT      string    `json:"Msg"`
	useremail string    `json:"email"`
}
type MessageResponse struct {
	CreatedAt time.Time `json:"createdAt"`
	TEXT      string    `json:"Msg"`
	useremail string    `json:"email"`
}

func createMsgTable() error {
	query := `
	CREATE TABLE Msg (
		CreatedAt timestamp NOT NULL,
		text text NOT NULL,
		useremail text NOT NULL,
		msgID serial PRIMARY KEY
	);
    `
	err := s.db.Exec(query)
	return err
	defer CloseDatabase()
}
func CloseDatabase() error {
	return db.Close()
}

//start.go
