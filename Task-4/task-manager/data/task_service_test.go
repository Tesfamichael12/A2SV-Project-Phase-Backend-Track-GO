package data

import (
	"task_manager/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


var originalTasks = []models.Task{
	{ID: 1, Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "In Progress"},
	{ID: 2, Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: 3, Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func resetTasks() {
	tasks = make([]models.Task, len(originalTasks))
	copy(tasks, originalTasks)
}

func TestGetAllTasks(t *testing.T) {
	resetTasks()

	result := GetAllTasks()

	assert.Equal(t, len(tasks), len(result))
}

func TestGetTaskByID(t *testing.T) {
	resetTasks()

	task, found := GetTaskByID(1)

	assert.True(t, found)
	assert.Equal(t, 1, task.ID)
}

func TestAddTask(t *testing.T) {
	resetTasks()

	newTask := models.Task{
		Title:       "New Task",
		Description: "A new task",
		DueDate:     time.Now().AddDate(0, 0, 3),
		Status:      "In Progress",
	}

	AddTask(newTask)

	if len(tasks) == 0 {
		t.Fatalf("tasks slice is empty after AddTask")
	}

	added := tasks[len(tasks)-1]
	assert.Equal(t, newTask.Title, added.Title)
	assert.Equal(t, newTask.Description, added.Description)
	assert.Equal(t, newTask.Status, added.Status)
}

func TestUpdateTask(t *testing.T) {
	resetTasks()

	updated := models.Task{
		Title:       "Updated Title",
		Description: "Updated Desc",
		DueDate:     time.Now().AddDate(0, 0, 4),
		Status:      "Completed",
	}

	ok := UpdateTask(1, updated)
	assert.True(t, ok)
	task, found := GetTaskByID(1)
	assert.True(t, found)
	assert.Equal(t, "Updated Title", task.Title)
}

func TestDeleteTask(t *testing.T) {
	resetTasks()

	ok := DeleteTask(1)
	assert.True(t, ok)
	
	_, found := GetTaskByID(1)
	assert.False(t, found)
}
