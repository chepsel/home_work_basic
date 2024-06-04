package source

import (
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB     *sqlx.DB
	Logger *slog.Logger
}

type SentinelError string

func (err SentinelError) Error() string {
	return string(err)
}

const (
	WrongParam     SentinelError = "wrong param"
	MissingID      SentinelError = "id is missing"
	NoRowsAffected SentinelError = "none rows affected"
	DatabaseIsDown SentinelError = "database is down"
)

func (src *Database) PingDS() error {
	var i int8
	queryType := "PingDataSource"
	err := src.DB.Ping()
	if err != nil {
		for i < 10 && err != nil {
			src.LogError(queryType, err)
			time.Sleep(5 * time.Second)
			err = src.DB.Ping()
			i++
		}
		return err
	}
	go src.LogDBResult(queryType, "database is up")
	return nil
}

func (src *Database) Connect(db string, dsn string) (err error) {
	var i int8
	queryType := "Connection to DB"
	src.DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		for i < 10 && err != nil {
			src.LogError(queryType, err)
			time.Sleep(5 * time.Second)
			src.DB, err = sqlx.Connect("postgres", dsn)
			i++
		}
		return err
	}
	go src.LogDBResult(queryType, "connection is ok")
	return nil
}
