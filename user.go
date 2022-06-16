package filemanager

type User struct {
	Uid      string `json:"-" db:"uid"`
	Email    string `json:"email" db:"email" binding:"email" `
	Password string `json:"password" db:"password_hash"`
}
