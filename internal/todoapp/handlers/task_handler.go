package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/todo-app-go/internal/todoapp/task"
)

func CreateTaskHandler(r task.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var task task.Task
		err := ctx.BindJSON(&task)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		res, err := r.CreateTask(task)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func GetTasks(r task.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tasks, err := r.Get()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, tasks)
	}
}

func GetTaskByID(r task.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		task, err := r.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, task)
	}
}

func UpdateTask(r task.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		completeTask, err := r.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		err = ctx.BindJSON(&completeTask)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		res, err := r.UpdateByID(id, completeTask.Title, completeTask.Description, completeTask.FinishDate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func DeleteTask(r task.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		deletedID, err := r.DeleteByID(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, deletedID)
	}
}
