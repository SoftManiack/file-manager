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
	GetItems(rootuid string) ([]item.Directory, []item.File, error)
	CreateDirectory(input item.NewDirectory, userUid string) (item.Directory, error)
	CreateTextFile(input item.NewFile) (item.File, error)
	Rename(uid string, newName item.Rename) error
	SetPath(uid, path, typeItem string) error
	SetSize(uid, size, typeItem string) error
	//Delete(uid string) error
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
