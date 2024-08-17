package repository

import (
	"errors"
	user "file-manager/dto"
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

func (r *AuthPostgres) DeleteUser(uidUser string) error {

	queryDeleteUser := fmt.Sprintf("DELETE FROM %s WHERE uid = $1", usersTable)

	_, err := r.db.Exec(queryDeleteUser, uidUser)

	return err
}

func (r *AuthPostgres) CreateUser(user user.User) (string, error) {
	fmt.Println("Create User")
	fmt.Println(user)

	tx, err := r.db.Begin()
	if err != nil {
		logrus.Errorln(err)
		return "", err
	}

	var uid string

	queryGetEmail := fmt.Sprintf("SELECT uid FROM %s WHERE email = $1", usersTable)
	queryCreateUser := fmt.Sprintf("INSERT INTO %s ( email, name, password_hash) values ($1, $2, $3) RETURNING uid", usersTable)
	queryRoot := fmt.Sprintf("INSERT INTO %s ( uid, users_uid, root_uid, name, is_favorites, count_element, path ) VALUES ($1, $2, $3, $4, $5, $6, $7)", directoriesTable)

	if err := r.db.Get(&uid, queryGetEmail, user.Email); err != nil {
		row := r.db.QueryRow(queryCreateUser, user.Email, user.Name, user.Password)

		fmt.Println("createUser")
		if err := row.Scan(&uid); err != nil {
			tx.Rollback()
			return "", err
		}

		fmt.Println("createRoot")
		_, err := r.db.Exec(queryRoot, uid, uid, uid, "Облачный диск", false, 0, "Облачный диск")
		if err != nil {
			tx.Rollback()
			return "", err
		}

	} else {
		tx.Rollback()
		logrus.Errorln(err)
		return "", errors.New("такой email уже существет")
	}

	return uid, tx.Commit()
}

func (r *AuthPostgres) GetUser(email, password string) (user.User, error) {

	var user user.User

	query := fmt.Sprintf("SELECT uid FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	user.Email = email

	fmt.Println(err)
	if err != nil && err.Error() == "sql: no rows in result set" {
		err = errors.New("неверная электронная почта или пароль")
	}

	fmt.Println("user")
	fmt.Println(user)

	return user, err
}
