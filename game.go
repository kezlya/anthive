package anthive

type Game struct {
	Id       string `json:"id"`
	Author   string `json:"author"`
	Flavor   string `json:"flavor"`
	Version  string `json:"version"`
	Limit    uint   `json:"limit"`
	Age      uint   `json:"age"`
	Bots     []Bot  `json:"bots"`
	Map      Map    `json:"map"`
	Goal     Goal   `json:"goal"`
	Queued   uint   `json:"queued"`
	Started  uint   `json:"started"`
	Finished uint   `json:"finished"`
	Duration uint   `json:"duration"`
}

type Goal struct {
	Age    uint    `json:"age"`
	Ants   uint    `json:"ants"`
	Hive   uint    `json:"hive"`
	Alone  bool    `json:"alone"`
	Points []Point `json:"points"`
}
