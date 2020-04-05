package anthive

type BotRequest struct {
	ID      string          `json:"id"`
	Tick    uint16          `json:"tick"`
	Ants    map[uint16]*Ant `json:"ants"`
	*Canvas `json:"canvas"`
}

type Ant struct {
	Event   string `json:"event"`
	Errors  int    `json:"errors"`
	Age     uint8  `json:"age"`
	Health  uint8  `json:"health"`
	Payload uint8  `json:"payload"`
	X       uint8  `json:"x"`
	Y       uint8  `json:"y"`
}

type Canvas struct {
	Width  uint8     `json:"width"`
	Height uint8     `json:"height"`
	Cells  [][]*Cell `json:"cells"`
}
