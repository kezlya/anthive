package anthive

type Map struct {
	FirstFrame Frame
	Settings   MapSettings
}

type Frame struct {
}

type MapSettings struct {
	Width, Height, Food,
	Hunger, Skin uint8
	Dark bool
}

type Game struct {
	Queued, Started, Finished  int64
	ID, Author, Flavor, Status string
	Limit                      uint16
	Bots                       []Bot
	Map                        Map
}

type Sandbox struct {
	Code   string `json:"code" firestore:"code"`
	lang   string `json:"lang" firestore:"lang"`
	Result string `json:"result" firestore:"result"`
	Status string `json:"status" firestore:"status"`
	Queued int64  `json:"queued" firestore:"queued"`
}
