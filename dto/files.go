package dto

type NewDefaultFile struct {
	RootUid   string `json:"rootUid"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

type NewTextFile struct {
	RootUid string `json:"rootUid"`
	Name    string `json:"name"`
	Text    string `json:"text"`
}

type NewImageFile struct {
	RootUid string `json:"rootUid"`
	Name    string `json:"name"`
	Data    []byte `json:"data"`
}

type UpdateFile struct {
	Uid         string `json:"uid"`
	Name        string `json:"name"`
	IsFavorites bool   `json:"isFavorites"`
}

type UploadFile struct {
	RootUid   string `json:"rootUid"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
}

type File struct {
	Uid         string `json:"uid" db:"uid"`
	RootUid     string `json:"rootUid" db:"root_uid"`
	DateCreate  string `json:"dateCreate" db:"date_create"`
	DateUpdate  string `json:"dateUpdate" db:"date_update"`
	Name        string `json:"name" db:"name"`
	Path        string `json:"path" db:"path"`
	IsFavorites bool   `json:"isFavorite" db:"is_favorites"`
	IsLink      bool   `json:"IsLink" db:"is_link"`
	IsDelete    bool   `json:"IsDelete" db:"is_delete"`
	Size        int64  `json:"size" db:"size"`
	Type        string `json:"type" db:"type"`
	Data        string `json:"data"`
}
