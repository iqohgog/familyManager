package db

import (
	"database/sql"
	"fmt"
	"v1/familyManager/configs"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func New(conf *configs.Config) *Storage {
	storagePath := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		conf.Db.DatabaseUser,
		conf.Db.DatabasePassword,
		conf.Db.DatabasePort,
		conf.Db.DatabaseName,
	)
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		panic(err)
	}
	return &Storage{DB: db}
}
