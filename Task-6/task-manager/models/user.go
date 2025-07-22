package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User defines the structure for a user account.
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username" validate:"required,min=3"`
	Password string             `bson:"password" json:"password" validate:"required,min=6"`
}
