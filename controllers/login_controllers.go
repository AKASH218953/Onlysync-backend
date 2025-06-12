package controllers

import (
	"context"
	"net/http"
	"onlysync/configs"
	"onlysync/models"
	"onlysync/responces"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func Signup(c echo.Context) error {

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responces.LoginResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    map[string]interface{}{"error": err.Error()},
		})
	}
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, responces.LoginResponse{
			Status:  http.StatusBadRequest,
			Message: "Username, email, and password are required",
			Data:    map[string]interface{}{"error": "Missing required fields"},
		})
	}

	// Check if user already exists
	filters := bson.M{
		"email":    user.Email,
		"username": user.Username,
		"usermetadata": bson.M{
			"phone": user.UserMetadata.Phone,
		},
	}

	err := userCollection.FindOne(ctx, filters).Decode(&user)
	if err == nil {
		return c.JSON(http.StatusConflict, responces.LoginResponse{
			Status:  http.StatusConflict,
			Message: "User already exists",
			Data:    map[string]interface{}{"error": "User with this email already exists"},
		})

	} else if err != mongo.ErrNoDocuments {
		return c.JSON(http.StatusInternalServerError, responces.LoginResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error checking existing user",
			Data:    map[string]interface{}{"error": err.Error()},
		})
	}

	newuser := models.User{
		ID:        primitive.NewObjectID(),
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password, // In a real application, you should hash the password
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserMetadata: models.UserMetadat{
			Role:        user.UserMetadata.Role,
			Department:  user.UserMetadata.Department,
			Graducation: user.UserMetadata.Graducation,
			Location:    user.UserMetadata.Location,
			Phone:       user.UserMetadata.Phone,
			Active:      true,
		},
	}

	insertResult, err := userCollection.InsertOne(context.Background(), newuser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responces.LoginResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error inserting data",
			Data:    err.Error(),
		})
	}

	// Implement signup logic here
	return c.JSON(http.StatusAccepted, responces.LoginResponse{
		Status:  http.StatusAccepted,
		Message: "Signup successful",
		Data:    insertResult,
	})
}
