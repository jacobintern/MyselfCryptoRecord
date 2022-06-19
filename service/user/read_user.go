package user

import (
	"context"

	"github.com/jacobintern/MyselfCryptoRecord/service"
	"go.mongodb.org/mongo-driver/bson"
)

// UserReadModel is a model
type UserReadModel struct {
	Uid   string `bson:"_id,omitempty" josn:"uid"`
	Name  string `bson:"name" json:"userName"`
	Email string `bson:"email" json:"userEmail"`
}

func GetUserList() []*UserReadModel {
	collection := service.MongoDbContext(myAccList)

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []*UserReadModel
	if err = cursor.Decode(&results); err != nil {
		panic(err)
	}

	return results
}

func GetUserByAcc(acc string) (user *UserReadModel) {
	collection := service.MongoDbContext(myAccList)

	var result *UserReadModel
	err := collection.FindOne(context.TODO(), bson.M{"acc": acc}).Decode(result)
	if err != nil {
		panic(err)
	}

	return result
}
