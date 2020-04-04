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
