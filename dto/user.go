package dto

type User struct {
	Uid      string `json:"uid" db:"uid"`
	Email    string `json:"email" db:"email" binding:"email"`
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password_hash"`
}
