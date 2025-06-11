package playlist

type Playlist struct {
	Name       string  `json:"name"`
	Author     string  `json:"author"`
	Created_at string  `json:"created_at"`
	Tracks     []Track `json:"tracks"`
}

type Track struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Votes  int    `json:"votes"`
}

type UserService struct{}
