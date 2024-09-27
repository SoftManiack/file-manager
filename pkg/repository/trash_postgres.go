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

	var uidTrash string

	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	queryUpdateIsDelete := fmt.Sprintf("UPDATE %s SET is_delete = true WHERE uid = $1", filesTable)
	queryTrashUid := fmt.Sprintf("SELECT uid FROM %s WHERE users_uid = $1", trashTable)

	queryMoveTrash := fmt.Sprintf("INSERT INTO %s ( trash_uid, files_uid ) values ($1, $2)", trashFilesTable)

	_, err = r.db.Exec(queryUpdateIsDelete, uidFile)

	if err := r.db.Get(&uidTrash, queryTrashUid, uidUser); err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.db.Exec(queryMoveTrash, uidTrash, uidFile)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *TrashPostgres) RemoveTrashDirectory(uidDir string) error {

	return nil
}

func (r *TrashPostgres) RemoveTrashFile(uidFile string) error {

	queryDeleteFileForTrash := fmt.Sprintf("DELETE FROM %s WHERE files_uid = $1", trashFilesTable)
	queryUpdateFile := fmt.Sprintf("UPDATE %s SET is_delete = false WHERE uid = $1", filesTable)

	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.db.Exec(queryDeleteFileForTrash, uidFile)

	_, err = r.db.Exec(queryUpdateFile, uidFile)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *TrashPostgres) DeleteTrashFile(uidFile string) error {

	// удалить еше из корзины

	queryDeleteFile := fmt.Sprintf("DELETE FROM %s WHERE uid = $1", filesTable)
	queryDeleteFileTrash := fmt.Sprintf("DELETE FROM %s WHERE files_uid = $1", trashFilesTable)

	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.db.Exec(queryDeleteFile, uidFile)

	_, err = r.db.Exec(queryDeleteFileTrash, uidFile)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *TrashPostgres) DeleteTrashDirectory(uid string) error {

	return nil
}

func (r *TrashPostgres) GetTrash(uidUser string) ([]directories.Directories, []files.File, error) {

	var getDrectories []directories.Directories
	var getFiles []files.File

	queryGetFiles := fmt.Sprintf("SELECT * FROM %s WHERE uid = ( SELECT files_uid FROM %s WHERE trash_uid = ( SELECT uid FROM %s WHERE users_uid = $1 ))", filesTable, trashFilesTable, trashTable)

	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return getDrectories, getFiles, err
	}

	if err := r.db.Select(&getFiles, queryGetFiles, uidUser); err != nil {
		tx.Rollback()
		return getDrectories, getFiles, err
	}

	return getDrectories, getFiles, tx.Commit()
}
