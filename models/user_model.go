package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id"`
	Username     string             `json:"username"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	FirstName    string             `json:"first_name"`
	LastName     string             `json:"last_name"`
	ProfilePic   string             `json:"profile_pic"`
	Active       bool               `json:"active"`
	Status       bool               `json:"status"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserMetadata UserMetadat        `json:"usermatadata"`
}

type UserMetadat struct {
	Role        string `json:"role"`
	Department  string `json:"department"`
	Graducation string `json:"graduation"`
	Location    string `json:"location"`
	Phone       string `json:"phone"`
	Active      bool   `json:"active"`
}
