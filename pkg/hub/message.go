package hub

// WSEvent to frontend
type WSEvent struct {
	ChannelID string `json:"-"`
	Name      string `json:"name"`
	Event     string `json:"event"`
}

// WSCommand from frontend
type WSCommand struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}
