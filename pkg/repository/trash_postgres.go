package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TrashPostgres struct {
	db *sqlx.DB
}

func NewTrashPostgres(db *sqlx.DB) *TrashPostgres {
	return &TrashPostgres{db: db}
}

func (r *TrashPostgres) DeleteDirectory(uidDir string) error {

	return nil
}

func (r *TrashPostgres) DeleteFile(uidFile string) error {

	// изменить файл
	// сохранить в file_trash

	query := fmt.Sprintf("UPDATE %s SET count_element = count_element + 1 WHERE uid = $1", directoriesTable)

	return nil
}
