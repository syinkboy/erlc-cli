package erlc

type ServerInfo struct {
	Name           string `json:"Name"`
	CurrentPlayers int    `json:"CurrentPlayers"`
	MaxPlayers     int    `json:"MaxPlayers"`
	JoinKey        string `json:"JoinKey"`

	Players []Player `json:"Players,omitempty"`
	Staff   *Staff   `json:"Staff,omitempty"`
	Queue   []int64  `json:"Queue,omitempty"`
}

type Player struct {
	Team        string `json:"Team"`
	Player      string `json:"Player"` 
	Permission  string `json:"Permission"`
	WantedStars int    `json:"WantedStars"`
}

type Staff struct {
	Admins  map[string]string `json:"Admins"`
	Mods    map[string]string `json:"Mods"`
	Helpers map[string]string `json:"Helpers"`
}