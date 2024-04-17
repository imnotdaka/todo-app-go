package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/todo-app-go/cmd/config"
	"github.com/imnotdaka/todo-app-go/internal/database"
	"github.com/imnotdaka/todo-app-go/internal/todoapp/handlers"
	"github.com/imnotdaka/todo-app-go/internal/todoapp/task"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	db, err := database.NewDB(cfg.DB)
	if err != nil {
		return err
	}

	app := gin.New()

	app.GET("/ping", handlers.Ping())
	app.GET("/tasks", handlers.GetTasks(task.NewRepository(db)))
	app.GET("/tasks/:id", handlers.GetTaskByID(task.NewRepository(db)))
	app.POST("/tasks", handlers.CreateTaskHandler(task.NewRepository(db)))
	app.PUT("/tasks/:id", handlers.UpdateTask(task.NewRepository(db)))
	app.DELETE("/tasks/:id", handlers.DeleteTask(task.NewRepository(db)))

	err = app.Run()
	if err != nil {
		return err
	}

	return nil

}
