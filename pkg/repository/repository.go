package repository

import (
	directories "file-manager/dto"
	files "file-manager/dto"
	user "file-manager/dto"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user user.User) (string, error)
	GetUser(username, password string) (user.User, error)
	DeleteUser(uidUser string) error
}

type Directories interface {
	GetDirectory(uidDir string) ([]directories.Directories, []files.File, error)
	CreateDirectory(userUid string, input directories.NewDirectory) (directories.Directories, error)
	UpdateDirectory(input directories.UpdateDirectory) (directories.Directories, error)
	DeleteDirectory(uidDir string) error
	CopyDirectory(uidDir, uidTargetRoot string) error
	MoveDirectory(uidDir, uidTargetRoot string) error
}

type Files interface {
	CreateFile(uid string, input files.NewFile) (files.File, error)
	UpdateFile(input files.UpdateFile) (files.File, error)
	DeleteFile(uidFile string) error
	CopyFile(uidFile, uidTargetRoot string) error
	MoveFile(uidFile, uidTargetRoot string) error
}

type Repository struct {
	Authorization
	Directories
	Files
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Directories:   NewDirectoriesPostgres(db),
		Files:         NewFilesPostgres(db),
	}
}
