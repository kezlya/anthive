package anthive

type Map struct {
	Init     [][]Cell `json:"init" firestore:"init"`
	Settings Settings `json:"settings" firestore:"settings"`
}

type Settings struct {
	Width  uint8 `json:"width" firestore:"width"`
	Height uint8 `json:"height" firestore:"height"`
	Food   uint8 `json:"food" firestore:"food"`
	Hunger uint8 `json:"hunger" firestore:"hunger"`
	Theme  uint8 `json:"theme" firestore:"theme"`
}

type Cell struct {
	Food    int    `json:"food,omitempty" firestore:"food"`
	Terrain string `json:"terrain,omitempty" firestore:"terrain"`
	Hive    string `json:"hive,omitempty" firestore:"hive"`
	Ant     string `json:"ant,omitempty" firestore:"ant"`
}
