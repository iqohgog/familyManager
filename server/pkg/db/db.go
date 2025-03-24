package db

import (
	"database/sql"
	"fmt"
	"os"
	"v1/familyManager/configs"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func New(conf *configs.Config) *Storage {
	host := os.Getenv("DATABASE_HOST")
	storagePath := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Db.DatabaseUser,
		conf.Db.DatabasePassword,
		host,
		conf.Db.DatabasePort,
		conf.Db.DatabaseName,
	)
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		panic(err)
	}
	return &Storage{DB: db}
}
