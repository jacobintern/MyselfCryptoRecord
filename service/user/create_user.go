package user

import (
	"context"
	"log"

	"github.com/jacobintern/MyselfCryptoRecord/service"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser is a model
type CreateUserModel struct {
	Uid      string `bson:"_id,omitempty" json:"uid"`
	Name     string `bson:"name" json:"userName"`
	Account  string `bson:"acc" json:"userAcc"`
	Password string `bson:"pswd" json:"userPswd"`
	Email    string `bson:"email" json:"userEmail"`
}

func UserCreate(model *CreateUserModel) *mongo.InsertOneResult {
	mongoDB := service.ConnectionInfo{
		DBName:         "crypto",
		CollectionName: "my_acc_list",
	}
	collection := service.MongoDbContext(mongoDB)

	res, insertErr := collection.InsertOne(context.Background(), model)

	if insertErr != nil {
		log.Fatal(insertErr)
	}

	return res
}
