package anthive

type MapSettings struct {
	Width, Height, Food,
	Hunger, Skin uint8
	Dark bool
}

type Engineer struct {
	repo, sha        string
	Span, Username   string
	Version, Picture string
	FC, Skin, X, Y   uint8
	Wealth           uint16
}
