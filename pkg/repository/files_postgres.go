package repository

import (
	"errors"
	files "file-manager/dto"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type FilesPostgres struct {
	db *sqlx.DB
}

func NewFilesPostgres(db *sqlx.DB) *FilesPostgres {
	return &FilesPostgres{db: db}
}

func (r *FilesPostgres) CreateFile(newFile files.NewFile) (files.File, error) {

	var file files.File
	var uidFile string
	var uid string = ""
	//var names []string

	tx, err := r.db.Begin()

	if err != nil {
		logrus.Errorln(err)
		return files.File{}, err
	}

	//queryGetFilesNameFromDir := fmt.Sprintf("SELECT name FROM %s WHERE root_uid = $1", filesTable)
	queryGetUidForName := fmt.Sprintf("SELECT uid FROM %s WHERE name = $1", filesTable)
	queryCreateFile := fmt.Sprintf("INSERT INTO %s ( root_uid, name, path, size, data ) VALUES ($1, $2, $3, $4, $5) RETURNING uid", filesTable)
	queryGetNewFile := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", filesTable)

	/*if err := r.db.Select(&names, queryGetFilesNameFromDir, newFile.RootUid); err != nil {
		tx.Rollback()
		return file, err
	}

	if len(names) > 0 {
		tx.Rollback()
		return file, errors.New("такое имя уже существует")
	}*/

	err = r.db.Get(&uid, queryGetUidForName, newFile.Name)

	if uid != "" {
		tx.Rollback()
		return file, errors.New("такое имя уже существует")
	}

	row := r.db.QueryRow(queryCreateFile, newFile.RootUid, newFile.Name, newFile.Path, newFile.Size, newFile.Data)

	if err := row.Scan(&uidFile); err != nil {
		tx.Rollback()
		return file, err
	}

	if err := r.db.Get(&file, queryGetNewFile, uidFile); err != nil {
		tx.Rollback()
		return file, err
	}

	return file, tx.Commit()
}

func (r *FilesPostgres) UpdateFile(input files.UpdateFile) (files.File, error) {

	var file files.File

	tx, err := r.db.Begin()

	queryUpdateFile := fmt.Sprintf("UPDATE %s SET name = $1, is_favorites = $2, date_update = DEFAULT WHERE uid = $3", filesTable)
	queryGetNewFile := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", filesTable)

	_, err = r.db.Exec(queryUpdateFile, input.Name, input.IsFavorites, input.Uid)

	if err != nil {
		tx.Rollback()
		return file, err
	}

	if err := r.db.Get(&file, queryGetNewFile, input.Uid); err != nil {
		tx.Rollback()
		return file, err
	}

	return file, tx.Commit()
}

func (r *FilesPostgres) CopyFile(uidFile, uidTargetRoot string) error {

	return nil
}

func (r *FilesPostgres) MoveFile(uidFile, uidTargetRoot string) error {

	return nil
}

func (r *FilesPostgres) UpdateTextFile(updateFile files.UpdateTextFile) error {

	tx, err := r.db.Begin()

	queryUpdateFile := fmt.Sprintf("UPDATE %s SET name = $1, is_favorites = $2, data = $3 WHERE uid = $4", filesTable)

	_, err = r.db.Exec(queryUpdateFile, updateFile.Name, updateFile.IsFavorites, []byte(updateFile.Text), updateFile.Uid)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

/*func (r *FilesPostgres) getPath(uidDir, name string) (string, error) {

	var path string
	var rootDirectory directories.Directories

	queryGetRoot := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", directoriesTable)

	if err := r.db.Get(&rootDirectory, queryGetRoot, uidDir); err != nil {
		return "", err
	}
	_, err = r.db.Exec(queryUpdateFile, input.Name, input.IsFavorites, input.Uid) + name

	return path, nil
}*/
