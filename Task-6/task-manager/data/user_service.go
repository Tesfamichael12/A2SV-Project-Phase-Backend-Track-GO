package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

// initUserCollection is called by InitDB to set up the user collection.
func initUserCollection(client *mongo.Client) {
	userCollection = client.Database("taskdb").Collection("users")
}

// HashPassword securely hashes a password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // 14 is a good cost factor
	return string(bytes), err
}

// CheckPasswordHash compares a plain-text password with a hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CreateUser adds a new user to the database after hashing their password.
func CreateUser(user models.User) error {
	// Check if a user with the same username already exists
	var existingUser models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return errors.New("user with this username already exists")
	}
	if err != mongo.ErrNoDocuments {
		return err // A different database error occurred
	}

	// Hash the password before storing it
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	_, err = userCollection.InsertOne(context.Background(), user)
	return err
}

// GetUserByUsername retrieves a user by their username.
func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return user, err
}
