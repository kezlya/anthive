package anthive

type Map struct {
	Canvas Canvas `json:"canvas"`
	Food   uint32 `json:"food"`
	Hunger uint32 `json:"hunger"`
	Theme  uint32 `json:"theme"`
}

type Canvas struct {
	Width  uint32   `json:"width"`
	Height uint32   `json:"height"`
	Cells  [][]Cell `json:"cells"`
}

type Cell struct {
	Food    uint32 `json:"food,omitempty"`
	Terrain string `json:"terrain,omitempty"`
	Hive    string `json:"hive,omitempty"`
	Ant     string `json:"ant,omitempty"`
}

type Point struct {
	X uint32 `json:"x"`
	Y uint32 `json:"y"`
}
