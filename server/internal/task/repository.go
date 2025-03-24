package task

import (
	"database/sql"
	"v1/familyManager/pkg/db"
)

type TaskRepository struct {
	Storage *db.Storage
}

func NewTaskRepository(storage *db.Storage) *TaskRepository {
	return &TaskRepository{
		Storage: storage,
	}
}

func (repo *TaskRepository) Create(task *Task) (*Task, error) {
	stmt, err := repo.Storage.DB.Prepare(`
    INSERT INTO tasks(
        name, description, assignee_id, priority, family_id, creator_id
    )
    VALUES($1, $2, $3, $4, $5, $6)
    RETURNING id;
    `)
	if err != nil {
		return nil, err
	}
	var taskID int
	err = stmt.QueryRow(task.Name, task.Description, task.AssigneeID, task.Priority, task.FamilyID, task.CreatorID).Scan(&taskID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	task.ID = string(taskID)
	return task, nil
}

func (repo *TaskRepository) UpdateStatus(taskID string, status string) error {
	stmt, err := repo.Storage.DB.Prepare(`
    UPDATE tasks
    SET status = $1
    WHERE id = $2;
    `)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(status, taskID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TaskRepository) GetTaskByFamilyID(familyID string) (*[]Task, error) {
	stmt, err := repo.Storage.DB.Prepare(`
        SELECT id, name, description, assignee_id, priority, creator_id
        FROM tasks
        WHERE family_id = $1 && deleted_at IS NULL
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(familyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Name, &task.Description, &task.AssigneeID, &task.Priority, &task.CreatorID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return &tasks, nil
}

func (repo *TaskRepository) DeleteTaskByID(id string) error {
	stmt, err := repo.Storage.DB.Prepare(`
    UPDATE tasks
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE id = $1;
    `)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
