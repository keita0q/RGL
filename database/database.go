package database

import "github.com/keita0q/himatch/model"

type Database interface {
	LoadUser(aID string) (*model.User, error)
	SaveUser(*model.User) error
	DeleteUser(aID string) error

	LoadAllSpareTimes() ([]model.SpareTime, error)
	LoadSpareTime(aID string) (*model.SpareTime, error)
	SaveSpareTime(aSpareTime *model.SpareTime) error
	DeleteSpareTime(aID string) error
}

type NotFoundError struct {
	message string
}

func NewNotFoundError(aMessage string) *NotFoundError {
	return &NotFoundError{message: aMessage}
}

func (aNotFoundError *NotFoundError) Error() string {
	return aNotFoundError.message
}
