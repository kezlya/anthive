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

type Engineer struct {
	Wealth            uint16
	Username, Picture string
	Skin              uint8
}

type Game struct {
	Queued, Started, Finished  int64
	ID, Author, Flavor, Status string
	Limit                      uint16
	Bots                       []Bot
	Map                        Map
}

type Bot struct {
	ImageName string
	FC, X, Y  uint8
	Owner     Engineer
}
