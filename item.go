package filemanager

const (
	FILE string = "File"
	DIR         = "Directory"
)

type Item struct {
	Uid         string `json:"uid" db:"uid"`
	RootUid     string `json:"rootUid" db:"root_uid"`
	Date        string `json:"date" db:"date"`
	Extension   string `json:"extension" db:"extension"`
	Name        string `json:"name" db:"name"`
	Type        string `json:"type" db:"type"`
	IsFavorites bool   `json:"isFavorite" db:"is_favorites"`
	Link        bool   `json:"link" db:"link"`
	Size        string `json:"size" db:"size"`
	Trash       bool   `json:"trash" db:"trash"`
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
	Name string `json:"name"`
}
