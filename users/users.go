package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Username string			`json:"username,omitempty" bson:"username,omitempty" form:"username"`
	Password string			`json:"password,omitempty" bson:"password,omitempty" form:"password"`
}
