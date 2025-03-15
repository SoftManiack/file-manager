package repository

import (
	"database/sql"
	"errors"
	directories "file-manager/dto"
	dto "file-manager/dto"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DirectoriesPostgres struct {
	db *sqlx.DB
}

func NewDirectoriesPostgres(db *sqlx.DB) *DirectoriesPostgres {
	return &DirectoriesPostgres{db: db}
}

func (r *DirectoriesPostgres) GetDirectory(uidDir string) ([]dto.Directories, []dto.File, error) {

	var directories []dto.Directories
	var directoriesFiles []dto.File

	queryGetDirectories := fmt.Sprintf("SELECT * FROM %s WHERE root_uid = $1 AND root_uid <> uid AND is_delete <> true", directoriesTable)
	queryGetDefaultFiles := fmt.Sprintf("SELECT * FROM %s WHERE root_uid = $1 AND is_delete <> true", filesTable)

	err := r.db.Select(&directories, queryGetDirectories, uidDir)

	if err != nil {
		return directories, directoriesFiles, err
	}

	err = r.db.Select(&directoriesFiles, queryGetDefaultFiles, uidDir)

	if err != nil {
		return directories, directoriesFiles, err
	}

	return directories, directoriesFiles, nil
}

func (r *DirectoriesPostgres) CreateDirectory(userUid string, input dto.NewDirectory) (dto.Directories, error) {

	var newDirectory dto.Directories
	var uid string = ""
	var path string = ""

	queryGetUid := fmt.Sprintf("SELECT uid FROM %s WHERE name = $1 AND root_uid = $2", directoriesTable)
	queryCreateDirectory := fmt.Sprintf("INSERT INTO %s ( users_uid, root_uid, name, is_favorites, count_element, path ) values ($1, $2, $3, $4, $5, $6) RETURNING uid", directoriesTable)
	queryGetDirectory := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", directoriesTable)
	queryAddCount := fmt.Sprintf("UPDATE %s SET count_element = count_element + 1 WHERE uid = $1", directoriesTable)

	tx, err := r.db.Begin()
	if err != nil {
		return newDirectory, err
	}

	err = r.db.Get(&uid, queryGetUid, input.Name, input.RootUid)

	if uid != "" {
		tx.Rollback()
		return newDirectory, errors.New("такое имя уже существует")
	}

	if err == sql.ErrNoRows {

		path, err = r.getPath(input.RootUid, input.Name)

		if err != nil {
			return newDirectory, err
		}

		row := r.db.QueryRow(queryCreateDirectory, userUid, input.RootUid, input.Name, false, 0, path)
		if err := row.Scan(&uid); err != nil {
			tx.Rollback()
			return newDirectory, err
		}

		if err := r.db.Get(&newDirectory, queryGetDirectory, uid); err != nil {
			tx.Rollback()
			return newDirectory, err
		}

		_, err = r.db.Exec(queryAddCount, input.RootUid)

		if err != nil {
			tx.Rollback()
			return newDirectory, err
		}

	} else {
		tx.Rollback()
		return newDirectory, err
	}

	return newDirectory, tx.Commit()
}

func (r *DirectoriesPostgres) DownloadDirectory(uidDir string) (directories.DirectoriesTree, error) {

	var tree directories.DirectoriesTree

	// получить файлы
	// получить директории

	queryGetFiles := fmt.Sprintf("SELECT * FROM %s WHERE root_uid = $2", filesTable)
	queryGetDirectories := fmt.Sprintf("SELECT * FROM %s WHERE root_uid = $2", directoriesTable)

	err := r.db.Select(&tree.Files, queryGetFiles, uidDir)

	if err != nil {
		return directories.DirectoriesTree{}, err
	}

	err = r.db.Select(&tree.Directories, queryGetDirectories, uidDir)

	if err != nil {
		return directories.DirectoriesTree{}, err
	}

	for _, item := range tree.Directories {

		treeChild, err := r.getTree(item.Uid)

		if err != nil {
			return tree, nil
		}

		if len(treeChild.Directories) > 0 || len(treeChild.Files) > 0 {
			tree.DirectoriesTree = append(tree.DirectoriesTree, treeChild)
		}
	}

	// поиск диреткорий где чето лежит

	return tree, nil
}

func (r *DirectoriesPostgres) UpdateDirectory(input directories.UpdateDirectory) (directories.Directories, error) {

	var directory dto.Directories
	var uid string

	queryGetUid := fmt.Sprintf("SELECT uid FROM %s WHERE name = $1 AND root_uid = $2 AND uid != $3", directoriesTable)
	queryUpdateDirectory := fmt.Sprintf("UPDATE %s SET name = $1, is_favorites = $2, is_delete = $3 WHERE uid = $4", directoriesTable)
	queryGetDirectory := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", directoriesTable)

	fmt.Println(input)
	tx, err := r.db.Begin()
	if err != nil {
		return directory, err
	}

	err = r.db.Get(&uid, queryGetUid, input.Name, input.RootUid, input.Uid)

	if uid != "" {
		tx.Rollback()
		return directory, errors.New("такое имя уже существует")
	}

	_, err = r.db.Exec(queryUpdateDirectory, input.Name, input.IsFavorites, input.IsDelete, input.Uid)

	if err != nil {
		tx.Rollback()
		return directory, err
	}

	err = r.db.Get(&directory, queryGetDirectory, input.Uid)

	if err != nil {
		tx.Rollback()
		return directory, err
	}

	return directory, nil
}

func (r *DirectoriesPostgres) CopyDirectory(uidDir, uidTargetRoot string) error {

	return nil
}

func (r *DirectoriesPostgres) MoveDirectory(uidDir, uidTargetRoot string) error {

	return nil
}

func (r *DirectoriesPostgres) getPath(uidDir, name string) (string, error) {

	var path string
	var rootDirectory dto.Directories

	queryGetRoot := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", directoriesTable)

	if err := r.db.Get(&rootDirectory, queryGetRoot, uidDir); err != nil {
		fmt.Println(err)
		return "", err
	}

	path = rootDirectory.Path + "/" + name

	fmt.Println(path)

	return path, nil
}

func (r *DirectoriesPostgres) getTree(uidDir string) (directories.DirectoriesTree, error) {

	var tree directories.DirectoriesTree

	queryGetFiles := fmt.Sprintf("SELECT * FROM %s WHERE root_uid = $2", filesTable)
	queryGetDirectories := fmt.Sprintf("SELECT * FROM %s WHERE root_uid = $2", directoriesTable)

	err := r.db.Select(&tree.Files, queryGetFiles, uidDir)

	if err != nil {
		return directories.DirectoriesTree{}, err
	}

	err = r.db.Select(&tree.Directories, queryGetDirectories, uidDir)

	if err != nil {
		return directories.DirectoriesTree{}, err
	}

	for _, item := range tree.Directories {

		treeChild, _ := r.getTree(item.Uid)

		if len(treeChild.Directories) > 0 || len(treeChild.Files) > 0 {
			tree.DirectoriesTree = append(tree.DirectoriesTree, treeChild)
		}
	}

	// поиск диреткорий где чето лежит

	return tree, nil
}
