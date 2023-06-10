package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDocument struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	TransactionID string             `bson:"transactionid"`
	DocumentFile  string             `bson:"documentfile"`
	Name          string             `bson:"name"`
	Phone         string             `bson:"phone"`
	WalletAddress string             `bson:"walletaddress"`
	CreatedDate   time.Time          `bson:"created_date"`
}
