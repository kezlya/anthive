package anthive

type User struct {
}

type Engineer struct {
	Wealth            uint16
	Username, Picture string
	Skin              uint8
}

type Bot struct {
	ImageName string
	FC, X, Y  uint8
	Owner     Engineer
}
