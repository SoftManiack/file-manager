package repository

import (
	"database/sql"
	"errors"
	item "file-manager"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ItemsPostgres struct {
	db *sqlx.DB
}

func NewItemsPostgres(db *sqlx.DB) *ItemsPostgres {
	return &ItemsPostgres{db: db}
}

func (r *ItemsPostgres) GetItems(rootuid string) ([]item.Directory, []item.File, error) {

	var directories []item.Directory
	var files []item.File

	queryGetDirectories := fmt.Sprintf("SELECT * FROM %s WHERE root_uid = $1 AND root_uid <> uid", directoriesTable)
	err := r.db.Select(&directories, queryGetDirectories, rootuid)

	queryGetFiles := fmt.Sprintf("SELECT * FROM %s WHERE root_uid = $1 AND root_uid <> uid", filesTable)
	err = r.db.Select(&files, queryGetFiles, rootuid)

	return directories, files, err
}

func (r *ItemsPostgres) CreateDirectory(input item.NewDirectory, userUid string) (item.Directory, error) {

	var newDirectory item.Directory
	var uid string = ""

	queryGetUid := fmt.Sprintf("SELECT uid FROM %s WHERE name = $1", directoriesTable)
	queryCreateDirectory := fmt.Sprintf("INSERT INTO %s ( users_uid, root_uid, name, is_favorites, count_element) values ($1, $2, $3, $4, $5) RETURNING uid", directoriesTable)
	queryGetDirectory := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", directoriesTable)

	err := r.db.Get(&uid, queryGetUid, input.Name)

	if uid != "" {
		return newDirectory, errors.New("Такое имя уже существует")
	}

	fmt.Println(userUid)
	if err == sql.ErrNoRows {
		row := r.db.QueryRow(queryCreateDirectory, userUid, input.RootUid, input.Name, false, 0)
		if err := row.Scan(&uid); err != nil {
			return newDirectory, err
		}

		if err := r.db.Get(&newDirectory, queryGetDirectory, uid); err != nil {
			return newDirectory, err
		}

	} else {
		return newDirectory, err
	}

	return newDirectory, nil
}

func (r *ItemsPostgres) CreateTextFile(input item.NewFile) (item.File, error) {
	var newItem item.File
	var uid string = ""

	queryGetUid := fmt.Sprintf("SELECT uid FROM %s WHERE name = $1 AND root_uid = $2", filesTable)
	queryCreateFile := fmt.Sprintf("INSERT INTO %s ( root_uid, name, extension ) values ($1, $2, $3) RETURNING uid", filesTable)
	queryGetFile := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", filesTable)

	err := r.db.Get(&uid, queryGetUid, input.Name, input.RootUid)

	if uid != "" {
		return newItem, errors.New("Такое имя уже существует")
	}

	if err == sql.ErrNoRows {
		fmt.Println("ok")
		row := r.db.QueryRow(queryCreateFile, input.RootUid, input.Name, "txt")
		if err := row.Scan(&uid); err != nil {
			return newItem, err
		}

		if err := r.db.Get(&newItem, queryGetFile, uid); err != nil {
			return newItem, err
		}

	} else {
		return newItem, err
	}

	return newItem, nil

}

func (r *ItemsPostgres) Rename(uid string, input item.Rename) error {

	var queryRename string

	if input.Type == "directory" {
		queryRename = fmt.Sprintf("UPDATE %s SET name = $1 WHERE uid = $2", directoriesTable)
	} else {
		queryRename = fmt.Sprintf("UPDATE %s SET name = $1 WHERE uid = $2", filesTable)
	}

	_, err := r.db.Exec(queryRename, input.NewName, uid)

	return err
}

/*func (r *ItemsPostgres) Delete(uid string) error {

	queryDelete := fmt.Sprintf("DELETE FROM %s WHERE uid = $1", itemsTable)

	_, err := r.db.Exec(queryDelete, uid)

	return err
}*/
