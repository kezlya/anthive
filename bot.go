package anthive

type Bot struct {
	Source       string `json:"source" firestore:"source"`
	Version      string `json:"version" firestore:"version"`
	Conversion   uint8  `json:"conversion" firestore:"conversion"`
	*Coordinates `json:"coordinates" firestore:"coordinates"`
	*Engineer    `json:"conversion" firestore:"conversion"`
}

type Engineer struct {
	Score    uint16 `json:"score" firestore:"score"`
	Username string `json:"username" firestore:"username"`
	Avatar   string `json:"avatar" firestore:"avatar"`
	Skin     uint8  `json:"skin" firestore:"skin"`
}

type Coordinates struct {
	X uint8 `json:"x" firestore:"x"`
	Y uint8 `json:"y" firestore:"y"`
}
