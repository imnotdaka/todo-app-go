package task

import (
	"database/sql"
	"errors"

	"github.com/imnotdaka/todo-app-go/internal/database"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) CreateTask(task Task) (int64, error) {
	res, err := r.db.Exec(database.CreateTaskQuery, task.Title, task.Description, task.FinishDate)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (r Repository) Get() ([]Task, error) {
	rows, err := r.db.Query(database.GetAllQuery)
	if err != nil {
		return nil, err
	}
	tasks := []Task{}
	for rows.Next() {
		t := Task{}
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.FinishDate, &t.UpdateAt)
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r Repository) GetByID(id string) (Task, error) {
	row := r.db.QueryRow(database.GetByIDQuery, id)
	t := Task{}
	row.Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.FinishDate, &t.UpdateAt)
	return t, nil
}

func (r Repository) UpdateByID(id string, title string, description string, finishdate any) (any, error) {
	res, err := r.db.Exec(database.UpdateTaskByIdQuery, title, description, finishdate, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r Repository) DeleteByID(id string) (string, error) {
	res, err := r.db.Exec(database.DeleteByIDQuery, id)
	if err != nil {
		return "", err
	}
	affRows, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if affRows == 0 {
		return "", errors.New("no rows deleted")
	}

	return id, nil
}
