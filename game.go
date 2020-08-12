package anthive

type Game struct {
	Id       string `json:"id"`
	Author   string `json:"author"`
	Flavor   string `json:"flavor"`
	Version  string `json:"version"`
	Limit    int    `json:"limit"`
	Age      int    `json:"age"`
	Bots     []Bot  `json:"bots"`
	Map      Map    `json:"map"`
	Goal     Goal   `json:"goal"`
	Started  int    `json:"started"`
	Finished int    `json:"finished"`
}

type GameFirebase struct {
	Game       string `json:"game"`
	FirstFrame string `json:"firstFrame"`
}

type GameVideo struct {
}

type GameIndex struct {
}

// 	Queued   int `json:"queued"` Firebase and index

type Goal struct {
	Age    int     `json:"age"`
	Ants   int     `json:"ants"`
	Hive   int     `json:"hive"`
	Alone  bool    `json:"alone"`
	Points []Point `json:"points"`
}
