package anthive

type Map struct {
	Canvas Canvas `json:"canvas"`
	Food   uint   `json:"food"`
	Hunger uint   `json:"hunger"`
	Theme  uint   `json:"theme"`
}

type Canvas struct {
	Width  uint     `json:"width"`
	Height uint     `json:"height"`
	Cells  [][]Cell `json:"cells"`
}

type Cell struct {
	Food    uint   `json:"food,omitempty"`
	Terrain string `json:"terrain,omitempty"`
	Hive    string `json:"hive,omitempty"`
	Ant     string `json:"ant,omitempty"`
}

type Point struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}
