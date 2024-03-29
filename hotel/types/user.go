package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 7
)

type UpdateUserParams struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (p UpdateUserParams) ToBSON() bson.M {
	m := bson.M{}
	if len(p.FirstName) > 0 {
		m["firstname"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["lastname"] = p.LastName
	}
	return m
}

type CreateUserParams struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstNameLen {
		errors["firstname"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errors["lastname"] = fmt.Sprintf("lastname length should be at least %d characters", minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}
	if !isEmailValidate(params.Email) {
		errors["email"] = "email is invalid"
	}
	return errors
}

func isEmailValidate(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPasswrod string             `bson:"EncryptedPassword" json:"-"`
}

func NewUserFromPase(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPasswrod: string(encpw),
	}, nil
}

type AuthParams struct {
	Emial    string `json:"email"`
	Password string `json:"password"`
}
