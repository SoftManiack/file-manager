package repository

import (
	directories "file-manager/dto"
	files "file-manager/dto"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type FilesPostgres struct {
	db *sqlx.DB
}

func NewFilesPostgres(db *sqlx.DB) *FilesPostgres {
	return &FilesPostgres{db: db}
}

func (r *FilesPostgres) CreateFile(uid string, input files.NewFile) (files.File, error) {

	return files.File{}, nil
}

func (r *FilesPostgres) UpdateFile(input files.UpdateFile) (files.File, error) {

	return files.File{}, nil
}

func (r *FilesPostgres) DeleteFile(uidFile string) error {

	return nil
}

func (r *FilesPostgres) CopyFile(uidFile, uidTargetRoot string) error {

	return nil
}

func (r *FilesPostgres) MoveFile(uidFile, uidTargetRoot string) error {

	return nil
}

/* func (r *FilesPostgres) CreateTextFile(userUid string, input dto.NewTextFile) (dto.TextFile, error) {

	var newFile dto.TextFile

	var path string = ""
	var uidCheck string = ""
	var uidFile string = ""

	queryGetUid := fmt.Sprintf("SELECT uid FROM %s WHERE name = $1 AND root_uid = $2", textFilesTable)
	queryCreateFile := fmt.Sprintf("INSERT INTO %s ( root_uid, name, is_favorites, path ) values ($1, $2, $3, $4) RETURNING uid", textFilesTable)
	queryGetFile := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", textFilesTable)
	queryAddCount := fmt.Sprintf("UPDATE  %s SET count_element = 1 + ( SELECT count_element FROM  %s WHERE uid = $1 ) WHERE uid = $1", directoriesTable, directoriesTable)

	tx, err := r.db.Begin()
	if err != nil {
		return newFile, err
	}

	err = r.db.Get(&uidCheck, queryGetUid, input.Name, input.RootUid)

	fmt.Println("check uid")
	fmt.Println(uidCheck)
	fmt.Println(err)

	if uidCheck != "" {
		tx.Rollback()
		return newFile, errors.New("такое имя уже существует")
	}

	if err == sql.ErrNoRows {

		path, err = r.getPath(input.RootUid, input.Name)

		if err != nil {
			return newFile, err
		}

		fmt.Println("queryCreateFile")
		row := r.db.QueryRow(queryCreateFile, input.RootUid, input.Name, false, path)
		if err := row.Scan(&uidFile); err != nil {
			tx.Rollback()
			return newFile, err
		}

		if err := r.db.Get(&newFile, queryGetFile, uidFile); err != nil {
			tx.Rollback()
			return newFile, err
		}

		_, err = r.db.Exec(queryAddCount, input.RootUid, directoriesTable, directoriesTable)

		if err != nil {
			tx.Rollback()
			return newFile, err
		}

	} else {
		tx.Rollback()
		return newFile, err
	}

	return newFile, tx.Commit()

} */

func (r *FilesPostgres) getPath(uidDir, name string) (string, error) {

	fmt.Println(uidDir)
	fmt.Println(name)

	var path string
	var rootDirectory directories.Directories

	queryGetRoot := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", directoriesTable)

	if err := r.db.Get(&rootDirectory, queryGetRoot, uidDir); err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Printf("%#v\n", rootDirectory)

	fmt.Println("rootDirectory")
	fmt.Println(rootDirectory)
	path = rootDirectory.Path + "/" + name

	fmt.Println(path)

	return path, nil
}
