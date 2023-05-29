package handler

import (
	"context"

	db "github.com/The-Origin-Labs/landate/document/db"
	model "github.com/The-Origin-Labs/landate/document/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Init(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to Docuement Microserice ⚡",
	})
}

func AddUserDocs(c *fiber.Ctx) error {

	collection := db.GetCollection(db.MongoClient, "userdocs")

	var userDocs model.UserDocument
	if err := c.BodyParser(&userDocs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid Request",
			"error":   err.Error(),
		})
	}

	ctx := context.TODO()
	result, err := collection.InsertOne(ctx, userDocs)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid Request",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "User Document Inserted Successfully",
		"id":      result.InsertedID,
	})
}

func RemoveUserDocs(c *fiber.Ctx) error {

	id := c.Params("id")

	collection := db.GetCollection(db.MongoClient, "userdocs")
	ctx := context.TODO()
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid Request",
			"error":   err.Error(),
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Invalid Request",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusAccepted,
		"message": "Data removed successfully",
	})
}

func UpdateUserDocs(c *fiber.Ctx) error {

	id := c.Params("id")

	// var updatedData bson.M
	var updatedData model.UserDocument
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid Request",
			"error":   err.Error(),
		})
	}

	collection := db.GetCollection(db.MongoClient, "userdocs")
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedData}

	result, err := collection.UpdateByID(ctx, filter, update)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid Request",
			"error":   err.Error(),
		})
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Invalid Request",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "User(s) Data is Updated.",
	})
}

func GetUserDocsById(c *fiber.Ctx) error {

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}

	var userDocs model.UserDocument
	ctx := context.TODO()

	filter := bson.M{"_id": objectId}

	collection := db.GetCollection(db.MongoClient, "userdocs")
	err = collection.FindOne(ctx, filter).Decode(&userDocs)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": fiber.StatusAccepted,
		"data":   userDocs,
	})
}

func GetUserDocsByTxId(c *fiber.Ctx) error {

	var request struct {
		TxId string `json:"txid"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}

	ctx := context.TODO()

	filter := bson.M{"txid": request.TxId}

	collection := db.GetCollection(db.MongoClient, "userdocs")
	result, err := collection.Find(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}
	defer result.Close(ctx)

	var userDocs []model.UserDocument
	if err := result.All(ctx, &userDocs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}

	if len(userDocs) == 0 {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": fiber.StatusAccepted,
		"data":   userDocs,
	})
}

func GetUserDocsByPhone(c *fiber.Ctx) error {
	var request struct {
		PhoneNumber string `json:"phone"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}

	ctx := context.TODO()

	filter := bson.M{"phone": request.PhoneNumber}

	collection := db.GetCollection(db.MongoClient, "userdocs")
	result, err := collection.Find(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}
	defer result.Close(ctx)

	var userDocs []model.UserDocument
	if err := result.All(ctx, &userDocs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}

	if len(userDocs) == 0 {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": fiber.StatusAccepted,
		"data":   userDocs,
	})
}

func GetUserDocsByWalletAddress(c *fiber.Ctx) error {
	var request struct {
		WalletAddress string `json:"waller_address"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}

	ctx := context.TODO()

	filter := bson.M{"txid": request.WalletAddress}

	collection := db.GetCollection(db.MongoClient, "userdocs")
	result, err := collection.Find(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}
	defer result.Close(ctx)

	var userDocs []model.UserDocument
	if err := result.All(ctx, &userDocs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}

	if len(userDocs) == 0 {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": fiber.StatusAccepted,
		"data":   userDocs,
	})
}
