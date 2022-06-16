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

func (r *ItemsPostgres) GetItems(rootuid string) ([]item.Item, error) {

	var items []item.Item

	queryGetAll := fmt.Sprintf("SELECT * FROM %s WHERE root_uid = $1 AND root_uid <> uid", itemsTable)
	err := r.db.Select(&items, queryGetAll, rootuid)

	return items, err
}

func (r *ItemsPostgres) CreateDirectory(input item.NewDirectory) (item.Item, error) {

	var newItem item.Item
	var uid string = ""

	queryGetUid := fmt.Sprintf("SELECT uid FROM %s WHERE name = $1 AND root_uid = $2 AND type = $3", itemsTable)
	fmt.Println(queryGetUid)
	queryCreateDirectory := fmt.Sprintf("INSERT INTO %s ( root_uid, name, type, is_favorites, link ) values ($1, $2, $3, $4, $5) RETURNING uid", itemsTable)
	queryGetDirectory := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", itemsTable)

	err := r.db.Get(&uid, queryGetUid, input.Name, input.RootUid, item.DIR)

	if uid != "" {
		return newItem, errors.New("Такое имя уже существует")
	}

	if err == sql.ErrNoRows {
		row := r.db.QueryRow(queryCreateDirectory, input.RootUid, input.Name, item.DIR, false, false)
		if err := row.Scan(&uid); err != nil {
			return newItem, err
		}

		if err := r.db.Get(&newItem, queryGetDirectory, uid); err != nil {
			return newItem, err
		}

	} else {
		return newItem, err
	}

	return newItem, nil
}

func (r *ItemsPostgres) CreateTextFile(input item.NewFile) (item.Item, error) {
	var newItem item.Item
	var uid string = ""

	queryGetUid := fmt.Sprintf("SELECT uid FROM %s WHERE name = $1 AND root_uid = $2 AND type = $3", itemsTable)
	queryCreateFile := fmt.Sprintf("INSERT INTO %s ( root_uid, name, type, is_favorites, link ) values ($1, $2, $3, $4, $5) RETURNING uid", itemsTable)
	queryGetFile := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", itemsTable)

	err := r.db.Get(&uid, queryGetUid, input.Name, input.RootUid, item.FILE)

	if uid != "" {
		return newItem, errors.New("Такое имя уже существует")
	}

	if err == sql.ErrNoRows {
		fmt.Println("ok")
		row := r.db.QueryRow(queryCreateFile, input.RootUid, input.Name, item.FILE, false, false)
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

func (r *ItemsPostgres) Rename(uid, newName string) error {

	queryRename := fmt.Sprintf("UPDATE %s SET name = $1 WHERE uid = $2", itemsTable)

	_, err := r.db.Exec(queryRename, newName, uid)

	return err
}

func (r *ItemsPostgres) Delete(uid string) error {

	queryDelete := fmt.Sprintf("DELETE FROM %s WHERE uid = $1", itemsTable)

	_, err := r.db.Exec(queryDelete, uid)

	return err
}
