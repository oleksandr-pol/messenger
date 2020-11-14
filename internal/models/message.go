package models

import (
	"database/sql"
)

type MessageRepository interface {
	AllMessages() ([]*Message, error)
	CreateMessage(Message) (int, error)
	UpdateMessage(Message) error
	DeleteMessage(int) error
}

type Message struct {
	Id      int    `json:"message_id,omitempty"`
	UserId  int    `json:"user_id"`
	RoomId  int    `json:"room_id"`
	Message string `json:"message"`
}

type messageDB struct {
	*sql.DB
}

func (db *messageDB) AllMessages() ([]*Message, error) {
	rows, err := db.Query("select message_id, user_id, room_id, message from message")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make([]*Message, 0)
	for rows.Next() {
		message := new(Message)
		err := rows.Scan(&message.Id, &message.UserId, &message.RoomId, &message.Message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}

func (db *messageDB) CreateMessage(m Message) (int, error) {
	var id int

	sqlInsert := `
	INSERT INTO message (user_id, room_id, message)
	VALUES ($1, $2, $3)
	RETURNING message_id`
	// use exec
	row := db.QueryRow(sqlInsert, m.UserId, m.RoomId, m.Message)
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *messageDB) UpdateMessage(m Message) error {
	sqlUpdate := `UPDATE room SET message=$1 WHERE id=$2`

	_, err :=
		db.Exec(sqlUpdate,
			m.Message, m.Id)

	return err
}

func (db *messageDB) DeleteMessage(id int) error {
	_, err := db.Exec("DELETE FROM message WHERE id=$1", id)

	return err
}
