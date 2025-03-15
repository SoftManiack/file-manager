package dto

type NewDirectory struct {
	RootUid string `json:"rootUid"`
	Name    string `json:"name"`
}

type OptionsDirectory struct {
	Uid string `json:"uid"`
}

type UpdateDirectory struct {
	Uid         string `json:"uid"`
	RootUid     string `json:"rootUid"`
	Name        string `json:"name"`
	IsFavorites bool   `json:"isFavorites"`
	IsDelete    bool   `json:"isDelete"`
}

type DeleteDirectory struct {
	Uid     string `json:"uid"`
	RootUid string `json:"rootUid"`
	Name    string `json:"name"`
}

type Directories struct {
	Uid          string `json:"uid" db:"uid"`
	UsersUid     string `json:"uidUsers" db:"users_uid"`
	RootUid      string `json:"uidRoot" db:"root_uid"`
	DateCreate   string `json:"dateCreate" db:"date_create"`
	DateUpdate   string `json:"dateUpdate" db:"date_update"`
	Name         string `json:"name" db:"name"`
	IsFavorites  bool   `json:"isFavorite" db:"is_favorites"`
	IsLink       bool   `json:"isLink" db:"is_link"`
	IsDelete     bool   `json:"isDelete" db:"is_delete"`
	Size         int64  `json:"size" db:"size"`
	Path         string `json:"path" db:"path"`
	CountElement int64  `json:"countElement" db:"count_element"`
	Type         string `json:"type" db:"type"`
}

type DirectoriesTree struct {
	DirectoriesTree []DirectoriesTree
	Directories     []Directories
	Files           []File
}
