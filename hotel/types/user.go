package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 7
)

type CreateUserParams struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() []string {
	errors := []string{}
	if len(params.FirstName) < minFirstNameLen {
		errors = append(errors, fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen))
	}
	if len(params.LastName) < minLastNameLen {
		errors = append(errors, fmt.Sprintf("lastname length should be at least %d characters", minLastNameLen))
	}
	if len(params.Password) < minPasswordLen {
		errors = append(errors, fmt.Sprintf("password length should be at least %d characters", minPasswordLen))
	}
	if !isEmailValidate(params.Email) {
		errors = append(errors, fmt.Sprintf("email is invalid"))
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
