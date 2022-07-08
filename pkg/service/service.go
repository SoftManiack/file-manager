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
	GetItems(rootuid string) ([]item.Directory, []item.File, error)
	CreateDirectory(input item.NewDirectory, uid string) (item.Directory, error)
	CreateTextFile(input item.NewFile, uid string) (item.File, error)
	Rename(uid string, input item.Rename) error
	//Delete(uid, userUid string) error
	//Upload(file *multipart.FileHeader, uid string) (item.Item, error)
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
