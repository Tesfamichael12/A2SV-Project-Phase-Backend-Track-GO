package data

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"task_manager/models"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file for testing: %v", err)
	}

	uri := os.Getenv("DATABASE_URL")
	if uri == "" {
		log.Fatal("DATABASE_URL not set in .env file")
	}

	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB for testing: %v", err)
	}

	testDBName := "taskdb_test"
	taskCollection = client.Database(testDBName).Collection("tasks")

	exitCode := m.Run()

	if err := client.Database(testDBName).Drop(context.Background()); err != nil {
		log.Printf("Failed to drop test database: %v", err)
	}
	client.Disconnect(context.Background())

	os.Exit(exitCode)
}

func setupTest(t *testing.T) {
	err := taskCollection.Drop(context.Background())
	assert.NoError(t, err)
}

func TestAddTask(t *testing.T) {
	setupTest(t)
	newTask := models.Task{
		Title:       "Test Task",
		Description: "A task for testing",
		DueDate:     time.Now(),
		Status:      "Pending",
	}

	insertedID, err := AddTask(newTask)
	assert.NoError(t, err)
	assert.NotEqual(t, primitive.NilObjectID, insertedID)

	fetchedTask, err := GetTaskByID(insertedID)
	assert.NoError(t, err)
	assert.Equal(t, newTask.Title, fetchedTask.Title)
	assert.Equal(t, newTask.Description, fetchedTask.Description)
}

func TestGetTaskByID(t *testing.T) {
	setupTest(t)
	newTask := models.Task{Title: "Find Me"}
	insertedID, _ := AddTask(newTask)

	task, err := GetTaskByID(insertedID)
	assert.NoError(t, err)
	assert.Equal(t, insertedID, task.ID)
	assert.Equal(t, "Find Me", task.Title)

	nonExistentID := primitive.NewObjectID()
	_, err = GetTaskByID(nonExistentID)
	assert.Error(t, err)
	assert.Equal(t, mongo.ErrNoDocuments, err)
}

func TestGetAllTasks(t *testing.T) {
	setupTest(t)
	AddTask(models.Task{Title: "Task 1"})
	AddTask(models.Task{Title: "Task 2"})

	tasks, err := GetAllTasks()
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
}

func TestUpdateTask(t *testing.T) {
	setupTest(t)
	newTask := models.Task{Title: "Original Title"}
	insertedID, _ := AddTask(newTask)

	updatedTask := models.Task{
		Title:  "Updated Title",
		Status: "Completed",
	}

	modifiedCount, err := UpdateTask(insertedID, updatedTask)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), modifiedCount)

	fetchedTask, _ := GetTaskByID(insertedID)
	assert.Equal(t, "Updated Title", fetchedTask.Title)
	assert.Equal(t, "Completed", fetchedTask.Status)
}

func TestDeleteTask(t *testing.T) {
	setupTest(t)
	newTask := models.Task{Title: "To Be Deleted"}
	insertedID, _ := AddTask(newTask)

	deletedCount, err := DeleteTask(insertedID)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), deletedCount)

	_, err = GetTaskByID(insertedID)
	assert.Error(t, err)
	assert.Equal(t, mongo.ErrNoDocuments, err)
}
