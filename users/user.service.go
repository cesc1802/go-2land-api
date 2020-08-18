package users

import (
	"context"
	"fmt"
	"go-rest-api/db"
	"go.mongodb.org/mongo-driver/bson"
)

//need to be inject some service into service here
type UserService struct {
}

func NewUserServicce() *UserService {
	return &UserService{}
}

func (us UserService) FindByUserName(username string) (User, error) {
	collection := db.GetConnection().Database("2land").Collection("users")

	var u User
	filter := bson.D{{"username", username}}
	if err := collection.FindOne(context.TODO(), filter).Decode(&u); err != nil {
		fmt.Println(err.Error())
	}
	return u, nil
}

func FindByUsername(username string) User {
	collection := db.GetConnection().Database("2land").Collection("users")

	var u User
	filter := bson.D{{"username", username}}
	if err := collection.FindOne(context.TODO(), filter).Decode(&u); err != nil {
		fmt.Println(err.Error())
	}
	return u
}
