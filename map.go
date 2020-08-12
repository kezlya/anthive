package anthive

type Map struct {
	Init   string `json:"string"` //???
	Food   int    `json:"food"`
	Hunger int    `json:"hunger"`
	Theme  int    `json:"theme"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}
