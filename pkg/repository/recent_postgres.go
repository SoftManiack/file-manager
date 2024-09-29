package repository

import (
	files "file-manager/dto"

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
