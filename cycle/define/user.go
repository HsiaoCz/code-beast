package define

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"-"`
	Email    string             `bson:"email" json:"email"`
}

type CreateRequestParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
