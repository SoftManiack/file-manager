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

func (r *RecentPostgres) UpdateRecent(uidUser string) error {

	return nil
}

func (r *RecentPostgres) GetRecent(uidUser string) ([]files.File, error) {

	return []files.File{}, nil
}

func (r *RecentPostgres) SetRecent(uidUser, uidFile string) error {

	tx, err := r.db.Begin()

	querySetRecent := fmt.Sprintf("INSERT INTO %s ( users_uid, files_uid ) VALUES ($1, $2 )", recentTable)

	_, err = r.db.Exec(querySetRecent, uidUser, uidFile)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
