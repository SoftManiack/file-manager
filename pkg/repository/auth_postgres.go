package repository

import (
	"errors"
	user "file-manager"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user user.User) (string, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	var uid string

	queryGetEmail := fmt.Sprintf("SELECT uid FROM %s WHERE email = $1", usersTable)
	queryCreateUser := fmt.Sprintf("INSERT INTO %s ( email, password_hash ) values ($1, $2) RETURNING uid", usersTable)
	queryRoot := fmt.Sprintf("INSERT INTO %s (  uid, root_uid, name, size, type, is_favorites, link) VALUES ($1, $2, $3, $4, $5, $6, $7)", itemsTable)

	if err := r.db.Get(&uid, queryGetEmail, user.Email); err != nil {
		row := r.db.QueryRow(queryCreateUser, user.Email, user.Password)

		if err := row.Scan(&uid); err != nil {
			tx.Rollback()
			return "", err
		}

		_, err := r.db.Exec(queryRoot, uid, uid, "Root", 0, "Directory", false, false)
		if err != nil {
			tx.Rollback()
			return "", err
		}

	} else {
		tx.Rollback()
		return "", errors.New("Такой email уже существет")
	}

	return uid, tx.Commit()
}

func (r *AuthPostgres) GetUser(email, password string) (user.User, error) {

	var user user.User
	fmt.Println(password)
	query := fmt.Sprintf("SELECT uid FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)
	logrus.Errorln(err)
	return user, err
}
