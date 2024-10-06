package user

import (
	"github.com/google/uuid"
)

type ID string

func newID() ID {
	return ID(uuid.NewString())
}

func (i ID) String() string {
	return string(i)
}


type Name string

func newName(name string) (result Name) {
	return Name(name)
}

func (n Name) String() (result string) {
	return string(n)
}

type User struct {
	ID     ID
	Name   Name
	Avatar Avatar
}

func (User) Entity() {}


func NewUser(name string) (user User) {
	return User{
		ID:   newID(),
		Name: newName(name),
	}
}
