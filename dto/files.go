package dto

type NewFile struct {
	Name    string
	Path    string
	Size    int64
	Data    []byte
	RootUid string
}

type NewTextFile struct {
	RootUid string `json:"rootUid"`
	Name    string `json:"name"`
	Text    string `json:"text"`
}

type UpdateTextFile struct {
	Uid         string `json:"uid"`
	RootUid     string `json:"rootUid"`
	Name        string `json:"name"`
	IsFavorites bool   `json:"isFavorites"`
	Text        string `json:"text"`
}

type NewTableFile struct {
	RootUid string `json:"rootUid"`
	Name    string `json:"name"`
	Rows    string `json:"rows"`
}

type UpdateFile struct {
	Uid         string `json:"uid"`
	Name        string `json:"name"`
	OldName     string `json:"oldName"`
	IsFavorites bool   `json:"isFavorites"`
	RootUid     string `json:"rootUid"`
}

type UploadFile struct {
	RootUid   string `json:"rootUid"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
}

type CopyFile struct {
	Uid        string `json:"uid"`
	RootUid    string `json:"rootUid"`
	Name       string `json:"name"`
	RootDirUid string `json:"rootDirUid"`
}

type DownloadFile struct {
	Uid     string `json:"uid"`
	UidRoot string `json:"rootUid"`
	Name    string `json:"name"`
}

type UidFile struct {
	Uid string `json:"uid"`
}

type DeleteFile struct {
	Uid     string `json:"uid" binding:"required"`
	RootUid string `json:"rootUid" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

type File struct {
	Uid         string `json:"uid" db:"uid"`
	RootUid     string `json:"rootUid" db:"root_uid"`
	DateCreate  string `json:"dateCreate" db:"date_create"`
	DateUpdate  string `json:"dateUpdate" db:"date_update"`
	Name        string `json:"name" db:"name"`
	Path        string `json:"path" db:"path"`
	IsFavorites bool   `json:"isFavorite" db:"is_favorites"`
	IsLink      bool   `json:"isLink" db:"is_link"`
	IsDelete    bool   `json:"isDelete" db:"is_delete"`
	Size        int64  `json:"size" db:"size"`
	Type        string `json:"type" db:"type"`
	Data        string `json:"data"`
}
