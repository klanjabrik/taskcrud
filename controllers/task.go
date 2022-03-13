// controllers/task.go

package controllers

import (
	"net/http"
	"os"
	"time"

	api "taskcrud/api"
	"taskcrud/helpers"
	"taskcrud/models"
	"taskcrud/utils/token"

	"github.com/gin-gonic/gin"
)

type CreateTaskInput struct {
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
	UserId     uint   `json:"user_id"`
}

type UpdateTaskInput struct {
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
	UserId     uint   `json:"user_id"`
}

// GET /tasks
// Get all tasks
func FindTasks(c *gin.Context) {

	var tasks []models.Task

	loggedinUserId, _ := token.ExtractTokenID(c)
	m := map[string]interface{}{"user_id": loggedinUserId}
	if err := models.GetAllTask(&tasks, m); err != nil {
		api.RespondError(c, http.StatusNotFound,
			helpers.GetFormattedMessage("empty_notfound_formatted", "User"))
		return
	}

	api.RespondSuccess(c, http.StatusOK, tasks)
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	deadline, err := time.Parse(os.Getenv("DATETIME_LAYOUT"), input.Deadline)
	if err != nil {
		api.RespondError(c, http.StatusBadRequest,
			helpers.GetMessage("datetime_not_valid"))
		return
	}

	loggedinUserId, _ := token.ExtractTokenID(c)

	// Create task
	task := models.Task{AssingedTo: input.AssingedTo, Task: input.Task, Deadline: deadline, UserId: loggedinUserId}

	if models.AddNewTask(&task) != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	api.RespondSuccess(c, http.StatusOK, task)
}

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	var task models.Task

	loggedinUserId, _ := token.ExtractTokenID(c)
	m := map[string]interface{}{"user_id": loggedinUserId, "id": c.Param("id")}

	if err := models.GetOneTask(&task, m); err != nil {
		api.RespondError(c, http.StatusNotFound,
			helpers.GetFormattedMessage("empty_notfound_formatted", "Task"))
		return
	}

	api.RespondSuccess(c, http.StatusOK, task)
}

// PATCH /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {

	// Get model if exist
	var task models.Task
	if err := models.GetOneTaskId(&task, c.Param("id")); err != nil {
		api.RespondError(c, http.StatusNotFound,
			helpers.GetFormattedMessage("empty_notfound_formatted", "Task"))
		return
	}

	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	deadline, err := time.Parse(os.Getenv("DATETIME_LAYOUT"), input.Deadline)
	if err != nil {
		api.RespondError(c, http.StatusBadRequest,
			helpers.GetMessage("datetime_not_valid"))
		return
	}

	var updatedInput models.Task
	updatedInput.Deadline = deadline
	updatedInput.AssingedTo = input.AssingedTo
	updatedInput.Task = input.Task

	if err := models.PutOneTask(&task, &updatedInput); err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
	}

	api.RespondSuccess(c, http.StatusOK, task)
}

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	var task models.Task

	if err := models.GetOneTaskId(&task, c.Param("id")); err != nil {
		api.RespondError(c, http.StatusNotFound,
			helpers.GetFormattedMessage("empty_notfound_formatted", "Task"))
		return
	}

	if err := models.DeleteTask(&task); err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
	}

	api.RespondSuccess(c, http.StatusOK, true)
}
