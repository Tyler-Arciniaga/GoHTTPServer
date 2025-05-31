package playlist

type Playlist struct {
	Name       string  `json:"name"`
	Author     string  `json:"author"`
	Created_at string  `json:"created_at"`
	Tracks     []Track `json:"tracks"`
}

type Track struct {
	Title  string
	Artist string
	Album  string
}
