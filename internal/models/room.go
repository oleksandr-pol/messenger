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
	Id        int    `json:"room_id,omitempty"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"private"`
}

type RoomDB struct {
	*sql.DB
}

func (db *RoomDB) AllRooms() ([]*Room, error) {
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

func (db *RoomDB) CreateRoom(r Room) (int, error) {
	var id int

	sqlInsert := `
	INSERT INTO room (name, private)
	VALUES ($1, $2)
	RETURNING room_id`

	row := db.QueryRow(sqlInsert, r.Name, r.IsPrivate)
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *RoomDB) RoomById(id int) (*Room, error) {
	query := `select room_id, name, private from room where room.room_id = $1`

	row := db.QueryRow(query, id)
	room := new(Room)
	err := row.Scan(&room.Id, &room.Name, &room.IsPrivate)

	if err != nil {
		return nil, err
	}

	return room, nil
}

func (db *RoomDB) UpdateRoom(r Room) error {
	sqlUpdate := `UPDATE room SET name=$1, private=$2 WHERE id=$3`

	_, err :=
		db.Exec(sqlUpdate,
			r.Name, r.IsPrivate, r.Id)

	return err
}

func (db *RoomDB) DeleteRoom(id int) error {
	_, err := db.Exec("DELETE FROM room WHERE id=$1", id)

	return err
}
