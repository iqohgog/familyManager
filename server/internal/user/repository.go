package user

import (
	"database/sql"
	"v1/familyManager/pkg/db"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	Storage *db.Storage
}

func NewUserRepository(storage *db.Storage) *UserRepository {
	return &UserRepository{
		Storage: storage,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	stmt, err := repo.Storage.DB.Prepare(`
	INSERT INTO users(
		first_name, last_name, email, hash_password
	)
	VALUES($1, $2, $3, $4)
	`)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.HashPass).Scan()
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	stmt, err := repo.Storage.DB.Prepare(`
		SELECT id, first_name, last_name, email, hash_password FROM users
		WHERE email = $1
	`)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(email)
	var user User
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.HashPass)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetByID(id string) (*User, error) {
	stmt, err := repo.Storage.DB.Prepare(`
		SELECT first_name, last_name, email, hash_password FROM users
		WHERE ID = $1
	`)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	var user User
	err = row.Scan(&user.FirstName, &user.LastName, &user.Email, &user.HashPass)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Нужно ли реализовывать Update и Put? или Delete?
