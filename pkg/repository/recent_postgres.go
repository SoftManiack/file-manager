package repository

import (
	files "file-manager/dto"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RecentPostgres struct {
	db *sqlx.DB
}

func NewRecentPostgres(db *sqlx.DB) *RecentPostgres {
	return &RecentPostgres{db: db}
}

func (r *RecentPostgres) UpdateRecent(uidUser, uidFile string) error {

	var uid string
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	querySelectFileRecent := fmt.Sprintf("SELECT uid FROM %s WHERE files_uid = $1", recentTable)

	queryDeleteRecent := fmt.Sprintf("DELETE FROM %s WHERE files_uid = $1", recentTable)
	queryCreateRecent := fmt.Sprintf("INSERT INTO %s ( users_uid, files_uid) values ($1, $2)", recentTable)

	err = r.db.Get(&uid, querySelectFileRecent, uidFile)

	if uid != "" {
		_, err = r.db.Exec(queryDeleteRecent, uidFile)

	} else {
		_, err = r.db.Exec(queryCreateRecent, uidUser, uidFile)
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *RecentPostgres) GetRecent(uidUser string) ([]files.File, error) {

	return []files.File{}, nil
}
