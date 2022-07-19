package filemanager

type Items struct {
	Directory []Directory
	File      []File
}

type Directory struct {
	Uid          string `json:"uid" db:"uid"`
	UsersUid     string `json:"uidUsers" db:"users_uid"`
	RootUid      string `json:"rootUid" db:"root_uid"`
	DateCreate   string `json:"date_create" db:"date_create"`
	DateUpdate   string `json:"date_update" db:"date_update"`
	Name         string `json:"name" db:"name"`
	IsFavorites  string `json:"isFavorite" db:"is_favorites"`
	Size         string `json:"size" db:"size"`
	Path         string `json:"path" db:"path"`
	Type         string `json:"type" db:"type"`
	CountElement string `json:"countElement" db:"count_element"`
}

type File struct {
	Uid         string `json:"uid" db:"uid"`
	RootUid     string `json:"rootUid" db:"root_uid"`
	DateCreate  string `json:"date_create" db:"date_create"`
	DateUpdate  string `json:"date_update" db:"date_update"`
	Name        string `json:"name" db:"name"`
	IsFavorites string `json:"isFavorite" db:"is_favorites"`
	Extension   string `json:"extension" db:"extension"`
	Path        string `json:"path" db:"path"`
	Type        string `json:"type" db:"type"`
	Size        string `json:"size" db:"size"`
}

type NewDirectory struct {
	RootUid string `json:"rootUid"`
	Name    string `json:"name"`
}

type NewFile struct {
	RootUid string `json:"rootUid"`
	Name    string `json:"name"`
	Text    string `json:"text"`
}

type Rename struct {
	NewName string `json:"newName"`
	Type    string `json:"type"`
}
