package controllers

import (
	"github.com/AbhishekCS3459/HR_MS/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEmployeeByID(c *fiber.Ctx) error {
	collection := mg.Db.Collection("employees")
	employee := new(models.Employee)
	employee.ID = c.Params("id")
	if employee.ID == "" {
		return c.Status(400).SendString("ID is required")
	}
	filter := bson.D{{Key: "_id", Value: employee.ID}}
	err := collection.FindOne(c.Context(), filter).Decode(employee)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(employee)
}
