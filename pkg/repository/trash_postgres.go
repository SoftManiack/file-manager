package repository

import (
	directories "file-manager/dto"
	files "file-manager/dto"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TrashPostgres struct {
	db *sqlx.DB
}

func NewTrashPostgres(db *sqlx.DB) *TrashPostgres {
	return &TrashPostgres{db: db}
}

func (r *TrashPostgres) MoveTrashDirectory(uidUser, uidDirectory string) error {

	return nil
}

func (r *TrashPostgres) MoveTrashFile(uidUser, uidFile string) error {

	// изменить файл
	// получить id корзины
	// сохранить в file_trash

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	queryUpdateIsDelete := fmt.Sprintf("UPDATE %s SET is_delete = true WHERE uid = $1", filesTable)
	queryMoveTrash := fmt.Sprintf("INSERT INTO %s ( trash_uid, uid_file ) values ($1) WHERE trash_uid = ( SELECT uid FROM trash WHERE users_uid = $2)", trashFilesTable)

	_, err = r.db.Exec(queryUpdateIsDelete, uidFile)

	_, err = r.db.Exec(queryMoveTrash, uidUser, uidFile)

	return tx.Commit()
}

func (r *TrashPostgres) RemoveTrashDirectory(uid string) error {

	return nil
}

func (r *TrashPostgres) RemoveTrashFile(uid string) error {

	return nil
}

func (r *TrashPostgres) DeleteTrashFile(uid string) error {

	return nil
}

func (r *TrashPostgres) DeleteTrashDirectory(uid string) error {

	return nil
}

func (r *TrashPostgres) GetTrash() ([]directories.Directories, []files.File, error) {

	return []directories.Directories{}, []files.File{}, nil
}
