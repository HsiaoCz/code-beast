package datastruction

type User struct {
	ID       string `bson:"_id,omitempty" json:"id,omitempty"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"-"`
	Email    string `bson:"email" json:"email"`
}
