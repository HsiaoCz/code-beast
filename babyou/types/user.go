package types

type User struct {
	ID       string `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"-"`
	Email    string `bson:"email" json:"email"`
}
