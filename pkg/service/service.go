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
	DownloadDirectory(uid string) error
	MoveTrashDirectory(uidDir string) error
}

type Files interface {
	CreateFile(newFile files.NewFile) (files.File, error)
	UpdateFile(updateFIle files.UpdateFile) (files.File, error)
	MoveTrashFile(uidUser, uidFile string) error
}

type Trash interface {
	GetTrash(uidUser string) ([]directories.Directories, []files.File, error)
	DeleteTrashFile(DeleteFile files.DeleteFile, uidUser string) error
	DeleteTrashDirectory(uidDir, uidUser string) error
	RemoveTrashFile(uidFile string) error
}

type Service struct {
	Authorization
	Directories
	Files
	Trash
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Directories:   NewDirectoriesService(repos.Directories),
		Files:         NewFilesService(repos.Files, repos.Trash),
		Trash:         NewTrashService(repos.Trash),
	}
}
