package controllers

import (
	"github.com/AbhishekCS3459/HR_MS/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEmployee(c *fiber.Ctx) error {
	collection := mg.Db.Collection("employees")
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	employee.ID = primitive.NewObjectID().Hex()
	result, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	// from here I am just implementing the part to check weather the record is created or not
	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)
	createdEmployee := &models.Employee{}
	createdRecord.Decode(createdEmployee)

	return c.Status(201).JSON(createdEmployee)
}
