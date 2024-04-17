package database

var (
	CreateTaskQuery = `
	INSERT INTO tasks(title, description, finish_date) VALUES (?, ?, ?)
	`
	GetAllQuery = `
	SELECT * FROM tasks
	`
	GetByIDQuery = `
	SELECT * FROM tasks WHERE id = ?
	`
	UpdateTaskByIdQuery = `
	UPDATE tasks SET title=?, description=?, finish_date=? WHERE id=?
	`
	DeleteByIDQuery = `
	DELETE FROM tasks WHERE id=?
	`
)
