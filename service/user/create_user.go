package user

import (
	"context"
	"errors"

	"github.com/jacobintern/MyselfCryptoRecord/service"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserCreateModel is a model
type UserCreateModel struct {
	Uid      string `bson:"_id,omitempty" josn:"uid"`
	Name     string `bson:"name" json:"userName"`
	Account  string `bson:"acc" json:"userAcc"`
	Password string `bson:"pswd" json:"userPswd"`
	Email    string `bson:"email" json:"userEmail"`
}

var myAccList = service.ConnectionInfo{
	DBName:         "crypto_db",
	CollectionName: "my_acc_list",
}

func UserCreate(model *UserCreateModel) (*mongo.InsertOneResult, error) {
	// model valid
	if len(model.Account) <= 0 || len(model.Name) <= 0 || len(model.Password) <= 0 {
		return nil, errors.New("account, name, password is required")
	}

	collection := service.MongoDbContext(myAccList)
	// hash pwd
	model.Password = service.Str2MD5Str(model.Password)

	res, insertErr := collection.InsertOne(context.Background(), model)

	if insertErr != nil {
		return nil, insertErr
	}

	return res, nil
}
