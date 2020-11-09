package models

import (
	"database/sql"
)

type RoomsRepository interface {
	AllRooms() ([]*Room, error)
	CreateRoom(Room) (int, error)
	RoomById(int) (*Room, error)
	UpdateRoom(Room) error
	DeleteRoom(int) error
}

type Room struct {
	Id        int    `json:"uer_id,omitempty"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"private"`
}

type roomDB struct {
	*sql.DB
}

func (db *roomDB) AllRooms() ([]*Room, error) {
	rows, err := db.Query("select room_id, name, private from room")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := make([]*Room, 0)
	for rows.Next() {
		room := new(Room)
		err := rows.Scan(&room.Id, &room.Name, &room.IsPrivate)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (db *roomDB) CreateRoom(r Room) (int, error) {
	var id int

	sqlInsert := `
	INSERT INTO room (name, private)
	VALUES ($1, $2)
	RETURNING id`

	row := db.QueryRow(sqlInsert, r.Name, r.IsPrivate)
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *roomDB) RoomById(id int) (*Room, error) {
	query := `select room_id, name, private from room where room.room_id = $1`

	row := db.QueryRow(query, id)
	room := new(Room)
	err := row.Scan(&room.Id, &room.Name, &room.IsPrivate)

	if err != nil {
		return nil, err
	}

	return room, nil
}

func (db *roomDB) UpdateRoom(r Room) error {
	sqlUpdate := `UPDATE room SET name=$1, private=$2 WHERE id=$3`

	_, err :=
		db.Exec(sqlUpdate,
			r.Name, r.IsPrivate, r.Id)

	return err
}

func (db *roomDB) DeleteRoom(id int) error {
	_, err := db.Exec("DELETE FROM room WHERE id=$1", id)

	return err
}