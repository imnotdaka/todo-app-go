package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imnotdaka/todo-app-go/cmd/config"
)

func NewDB(dbconf config.DB) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbconf.User, dbconf.Psw, dbconf.Ip, dbconf.Port, dbconf.Database))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
