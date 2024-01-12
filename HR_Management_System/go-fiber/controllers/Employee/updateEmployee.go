package controllers

import (
	"github.com/AbhishekCS3459/HR_MS/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateEmployee(c *fiber.Ctx) error {
	collection := mg.Db.Collection("employees")
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	employee.ID = c.Params("id")
	// fmt.Println(employee)
	filter := bson.D{{Key: "_id", Value: employee.ID}}
	updatedEmployee, err := collection.UpdateOne(c.Context(), filter, bson.D{{Key: "$set", Value: employee}})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(updatedEmployee)
}
