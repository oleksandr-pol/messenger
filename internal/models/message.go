package models

import (
	"database/sql"
)

type MessageRepository interface {
	AllMessages() ([]*Message, error)
	CreateMessage(Message) (int, error)
	UpdateMessage(Message) error
	DeleteMessage(int) error
	RoomMessages(roomId int) ([]*Message, error)
}

type Message struct {
	Id        int    `json:"message_id,omitempty"`
	UserId    int    `json:"user_id"`
	RoomId    int    `json:"room_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type MessageDB struct {
	*sql.DB
}

func (db *MessageDB) AllMessages() ([]*Message, error) {
	rows, err := db.Query("select message_id, user_id, room_id, message from message")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanMessagesRows(rows)
}

func (db *MessageDB) CreateMessage(m Message) (int, error) {
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

func (db *MessageDB) UpdateMessage(m Message) error {
	sqlUpdate := `UPDATE room SET message=$1 WHERE id=$2`

	_, err :=
		db.Exec(sqlUpdate,
			m.Message, m.Id)

	return err
}

func (db *MessageDB) DeleteMessage(id int) error {
	_, err := db.Exec("DELETE FROM message WHERE id=$1", id)

	return err
}

func (db *MessageDB) RoomMessages(roomId int) ([]*Message, error) {
	query := `
		select
			m.*
		from
			room r
		inner join users u on
			u.user_id = r.room_id
		inner join message m on
			m.room_id = r.room_id
		where
			r.room_id = $1;
	`
	rows, err := db.Query(query, roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanMessagesRows(rows)
}

func scanMessagesRows(rows *sql.Rows) ([]*Message, error) {
	messages := make([]*Message, 0)
	for rows.Next() {
		message := new(Message)
		err := rows.Scan(&message.Id, &message.UserId, &message.RoomId, &message.Message, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
