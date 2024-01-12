package controllers

import (
	"github.com/AbhishekCS3459/HR_MS/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var mg = &models.Mg

func GetEmployeesAll(c *fiber.Ctx) error {
	var employees []models.Employee = make([]models.Employee, 0)
	query := bson.D{{}}
	cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(employees)
}
