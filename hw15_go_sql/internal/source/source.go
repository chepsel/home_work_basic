package source

import (
	"log/slog"

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
)
