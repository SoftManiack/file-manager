package service

import (
	item "file-manager"
	user "file-manager"
	"file-manager/pkg/repository"
)

type Authorization interface {
	CreateUser(user user.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Items interface {
	GetItems(rootuid string) ([]item.Item, error)
	CreateDirectory(input item.NewDirectory, uid string) (item.Item, error)
	CreateTextFile(input item.NewFile, uid string) (item.Item, error)
	Rename(uid, newName string) error
	Delete(uid, userUid string) error
}

type Service struct {
	Authorization
	Items
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Items:         NewItemService(repos.Items),
	}
}
