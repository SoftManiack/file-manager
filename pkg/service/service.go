package service

import (
	directories "file-manager/dto"
	files "file-manager/dto"
	user "file-manager/dto"
	"file-manager/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user user.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
	DeleteUser(uidUser string) error
}

type Directories interface {
	GetDirectory(uid string) ([]directories.Directories, []files.File, error)
	CreateDirectory(uid string, input directories.NewDirectory) (directories.Directories, error)
	UpdateDirectory(input directories.UpdateDirectory) (directories.Directories, error)
}

type Files interface {
	UploadFile(rootUid, name string, size int64) (files.File, error)
}

type Service struct {
	Authorization
	Directories
	Files
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Directories:   NewDirectoriesService(repos.Directories),
		Files:         NewFilesService(repos.Files),
	}
}
