package repository

import (
	item "file-manager"
	user "file-manager"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user user.User) (string, error)
	GetUser(username, password string) (user.User, error)
}

type Items interface {
	GetItems(rootuid string) ([]item.Item, error)
	CreateDirectory(input item.NewDirectory) (item.Item, error)
	CreateTextFile(input item.NewFile) (item.Item, error)
	Rename(uid, newName string) error
	Delete(uid string) error
}

type Repository struct {
	Authorization
	Items
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Items:         NewItemsPostgres(db),
	}
}
