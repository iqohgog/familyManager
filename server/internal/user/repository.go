package user

import "v1/familyManager/pkg/db"

type UserRepository struct {
	Storage *db.Storage
}

func NewUserRepository(storage *db.Storage) *UserRepository {
	return &UserRepository{
		Storage: storage,
	}
}

// func (repo *UserRepository) Create(user *User) (*User, error) {
// 	result := repo.Database.DB.Create(user)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return user, nil
// }

// func (repo *UserRepository) GetByEmail(email string) (*User, error) {
// 	var user User
// 	result := repo.Database.DB.First(&user, "Email = ?", email)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &user, nil
// }
