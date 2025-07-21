package data

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task_manager/models"
)

var client *mongo.Client
var taskCollection *mongo.Collection

// InitDB initializes the MongoDB connection using the connection string from .env
func InitDB() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default connection string")
	}

	uri := os.Getenv("DATABASE_URL")
	if uri == "" {
		log.Fatal("DATABASE_URL not set in .env file")
	}

	var err error
	clientOptions := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	taskCollection = client.Database("taskdb").Collection("tasks")
	log.Println("Connected to MongoDB!")
	return nil
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := taskCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTaskByID retrieves a single task by its ID
func GetTaskByID(id primitive.ObjectID) (models.Task, error) {
	var task models.Task
	err := taskCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	return task, err
}

// AddTask adds a new task to the database
func AddTask(task models.Task) (primitive.ObjectID, error) {
	task.ID = primitive.NewObjectID()
	result, err := taskCollection.InsertOne(context.Background(), task)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// UpdateTask updates an existing task in the database
func UpdateTask(id primitive.ObjectID, updatedTask models.Task) (int64, error) {
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}
	result, err := taskCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

// DeleteTask deletes a task from the database
func DeleteTask(id primitive.ObjectID) (int64, error) {
	result, err := taskCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}