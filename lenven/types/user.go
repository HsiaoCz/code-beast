package types

import (
	"fmt"
	"regexp"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 7
)

type User struct {
	ID        string `bson:"_id" json:"id"`
	Firstname string `bson:"first_name" json:"first_name"`
	Lastname  string `bson:"last_name" json:"last_name"`
	Password  string `bson:"password" json:"password"`
	Email     string `bson:"email" json:"email"`
	IsAdmin   bool   `bson:"isAdmin" json:"isAdmin"`
}

type CreateUserParams struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.Firstname) < minFirstNameLen {
		errors["firstname"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen)
	}
	if len(params.Lastname) < minLastNameLen {
		errors["lastname"] = fmt.Sprintf("lastname length should be at least %d characters", minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}
	if !isEmailValidate(params.Email) {
		errors["email"] = fmt.Sprintf("email %s is invaild", params.Email)
	}
	return errors
}
func isEmailValidate(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
