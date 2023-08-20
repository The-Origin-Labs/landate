package handlers

import (
	"fmt"
	database "landate/authentication/db"
	models "landate/authentication/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Init(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"path":    ctx.Path(),
		"uri":     ctx.Request().URI().String(),
		"message": "Welcome to Landate Authentication SVC",
	})
}

// @Desp: Create User
// @Route:
// @Method: POST
func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
		"data":    user,
	})
}

// @Desp: Get User by Id
// @Route: /auth/user
// @Method: GET
func RetrieveUser(ctx *fiber.Ctx) error {
	user := models.User{}
	userId := ctx.Params("id")

	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "User not found",
		})
	}
	return ctx.Status(http.StatusOK).JSON(user)
}

// @Desp: Get all users
// @Route: /auth/users
// @Method: GET
func GetAllUsers(ctx *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Find(&users)
	return ctx.Status(http.StatusOK).JSON(users)
}

// @Desp: Delete user by id
// @Route: /auth/user/:id
// @Method: DELETE
func RemoveUser(ctx *fiber.Ctx) error {
	user := models.User{}
	userId := ctx.Params("id")

	result := database.DB.Where("id = ?", userId).First(user)
	if result.Error != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   result.Error.Error(),
			"message": "User not found",
		})
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("User(%s) deleted successfully", userId)})
}

// @Desp: Update User by Id
// @Route: /auth/user/:id
// @Method: UPDATE
func UpdateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := ctx.Params("id")
	var existingUser models.User
	if err := database.DB.First(&existingUser, userID).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Name = user.Name
	existingUser.WalletAddress = user.WalletAddress
	existingUser.Email = user.Email
	existingUser.UserPhoto = user.UserPhoto
	existingUser.Age = user.Age
	existingUser.Profession = user.Profession
	existingUser.Country = user.Country
	existingUser.City = user.City
	existingUser.PropertyOwned = user.PropertyOwned

	// Save the updated user data in the database
	if err := database.DB.Save(&existingUser).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update user",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user updated successfully",
		"data":    existingUser,
	})
}

// @Desp: Get User by Wallet Address
// @Route: /auth/user/:walletAddress
// @Method: GET
func GetUserByWalledId(ctx *fiber.Ctx) error {
	user := models.User{}
	walletAddress := ctx.Params("walletAddress")

	if err := database.DB.Where("walletAddress = ?", walletAddress).First(&user).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "User not found",
		})
	}
	return ctx.Status(http.StatusOK).JSON(user)
}
