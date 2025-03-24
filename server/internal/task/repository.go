package task

import "v1/familyManager/pkg/db"

type TaskRepository struct {
	Storage *db.Storage
}

func NewFamilyRepository(storage *db.Storage) *TaskRepository {
	return &TaskRepository{
		Storage: storage,
	}
}
