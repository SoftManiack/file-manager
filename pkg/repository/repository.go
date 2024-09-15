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
	CopyDirectory(uidDir, uidTargetRoot string) error
	MoveDirectory(uidDir, uidTargetRoot string) error
}

type Files interface {
	CreateFile(input files.NewFile) (files.File, error)
	UpdateFile(input files.UpdateFile) (files.File, error)
	CopyFile(uidFile, uidTargetRoot string) error
	MoveFile(uidFile, uidTargetRoot string) error
}

type Trash interface {
	MoveTrashDirectory(uidUser, uidDir string) error
	MoveTrashFile(uidUser, uidFile string) error
	RemoveTrashDirectory(uid string) error
	RemoveTrashFile(uid string) error
	DeleteTrashFile(uid string) error
	DeleteTrashDirectory(uid string) error
	GetTrash() ([]directories.Directories, []files.File, error)
}

type Repository struct {
	Authorization
	Directories
	Files
	Trash
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Directories:   NewDirectoriesPostgres(db),
		Files:         NewFilesPostgres(db),
		Trash:         NewTrashPostgres(db),
	}
}
