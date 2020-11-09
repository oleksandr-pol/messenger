package env

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	DbHostName string
	DbHostPort int
	DbUserName string
	DbPassword string
	DbName     string
}

// NewDb creates postgre db connection.
// It accepts db config as parameter.
// It returns db connection.
// It returns err if failed to connect to db.
func NewDb(c DbConfig) (*sql.DB, error) {
	pg_con_string := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.DbHostName, c.DbHostPort, c.DbUserName, c.DbPassword, c.DbName)
	db, err := sql.Open("postgres", pg_con_string)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
