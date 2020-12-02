package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

var DB *sql.DB

type User struct {
	Id   int
	Name string
}

func GetUserById(userID int) (u User, err error) {
	err = sql.ErrNoRows
	return u, errors.Wrap(err, "failed")
}
