package models

import "database/sql"

type ParticipantRepository interface {
	AddParticipant(Participant) (int, error)
	DeleteParticipant(int) error
}

type ParticipantDB struct {
	*sql.DB
}

type Participant struct {
	Id     int `json:"participant_id,omitempty"`
	UserId int `json:"user_id"`
	RoomId int `json:"room_id"`
}

func (db *ParticipantDB) AddParticipant(p Participant) (int, error) {
	var id int

	sqlInsert := `
	INSERT INTO participant (user_id, room_id)
	VALUES ($1, $2)
	RETURNING participant_id`

	row := db.QueryRow(sqlInsert, p.UserId, p.RoomId)
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *ParticipantDB) DeleteParticipant(id int) error {
	_, err := db.Exec("DELETE FROM participant WHERE id=$1", id)

	return err
}
