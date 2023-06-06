package handler

import (
	"net/http"
	database "github.com/The-Origin-Labs/landate/storage/database"
	models "github.com/The-Origin-Labs/landate/storage/models"

	"github.com/gofiber/fiber/v2"
)

// DESP: Base route request.
// METHOD: GET
func Init(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"path":    ctx.Path(),
		"uri":     ctx.Request().URI().String(),
		"message": "Welcome to Heimdal API 😍",
	})
}

// DESP: Add property and owner details.
// METHOD: POST
func AddProperty(ctx *fiber.Ctx) error {
	// property := new(models.Property)
	var property models.Property
	if err := ctx.BodyParser(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	// database.DB.Create(&property)
	// Create new records in the database
	if err := database.DB.Create(&property).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create property",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "property created successfully",
		"data":    property,
	})
}

// DESP: Get all property and owner details.
// METHOD: GET
func GetAllProperties(ctx *fiber.Ctx) error {
	properties := []models.Property{}
	database.DB.Find(&properties)
	return ctx.Status(http.StatusOK).JSON(properties)
}

// DESP: Get property and owner details.
// METHOD: GET
func GetProperty(ctx *fiber.Ctx) error {
	property := models.Property{}

	propertyID := ctx.Params("id")

	if err := database.DB.Where("id = ?", propertyID).First(&property).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
					"message" : "Property not found",
				})
	}

	return ctx.Status(http.StatusOK).JSON(property)
}

// DESP: Delete property and owner details.
// METHOD: DELETE
func DeleteProperty(ctx *fiber.Ctx) error {
	property := models.Property{}

	propID := ctx.Params("id") // Assuming the ID is passed as a route parameter

	// Query the database to find the property by ID
	result := database.DB.Where("id = ?", propID).First(&property)
	if result.Error != nil {
		// Property not found
		// Other database error
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
			"message": "Property not found",
		})
	}

	// Delete the property from the database
	if err := database.DB.Delete(&property).Error; err != nil {
		// Handle the delete error
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Property deleted successfully"})
}
