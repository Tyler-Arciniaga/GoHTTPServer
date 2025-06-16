package tracks

type Track struct {
	Title  string              `json:"title"`
	Artist string              `json:"artist"`
	Album  string              `json:"album"`
	Votes  int                 `json:"votes"`
	Voters map[string]struct{} `json:"voters"`
}
