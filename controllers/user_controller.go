package controllers

import (
	"context"
	"gofiber-app/database"
	"gofiber-app/models"
	"gofiber-app/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c *fiber.Ctx) error {
	var req models.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Check for missing fields
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing fields"})
	}

	// Check if email already exists in the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var existingUser models.User
	err := database.UserCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&existingUser)
	if err == nil {
		// Email already exists
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email is already in use"})
	}

	// Check if the same password is used twice (for security, passwords should be unique per user)
	existingPasswordCount, err := database.UserCollection.CountDocuments(ctx, bson.M{"password": req.Password})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error checking password uniqueness"})
	}
	if existingPasswordCount > 0 {
		// Password is already used by another user
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "This password has already been used"})
	}

	// Hash the password using the helper function
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	req.Password = hashedPassword

	// Insert the user into the database
	_, err = database.UserCollection.InsertOne(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
	var req models.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Check for missing fields
	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing fields"})
	}

	// Find the user by email
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user models.User
	err := database.UserCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		// Email not found in the database
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User doesn't exist"})
	}

	// Compare the provided password with the stored hashed password
	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{"message": "Login successful"})
}
