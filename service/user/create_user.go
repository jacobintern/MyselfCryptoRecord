package user

import (
	"context"
	"errors"

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

func UserCreate(model *CreateUserModel) (*mongo.InsertOneResult, error) {
	// model valid
	if len(model.Account) <= 0 || len(model.Name) <= 0 || len(model.Password) <= 0 {
		return nil, errors.New("account, name, password is required")
	}

	mongoDB := service.ConnectionInfo{
		DBName:         "crypto",
		CollectionName: "my_acc_list",
	}
	collection := service.MongoDbContext(mongoDB)

	res, insertErr := collection.InsertOne(context.Background(), model)

	if insertErr != nil {
		return nil, insertErr
	}

	return res, nil
}
