package filemanager

const (
	FILE string = "File"
	DIR         = "Directory"
)

type Item struct {
	Uid         string `json:"uid" db:"uid"`
	RootUid     string `json:"rootUid" db:"root_uid"`
	Name        string `json:"name" db:"name"`
	Size        string `json:"size" db:"size"`
	Type        string `json:"type" db:"type"`
	IsFavorites bool   `json:"isFavorite" db:"is_favorites"`
	Date        string `json:"date" db:"date"`
	Link        bool   `json:"link" db:"link"`
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
