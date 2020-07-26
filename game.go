package anthive

type Game struct {
	Id       string `json:"id"`
	Author   string `json:"author"`
	Flavor   string `json:"flavor"`
	Version  string `json:"version"`
	Limit    uint32 `json:"limit"`
	Age      uint32 `json:"age"`
	Bots     []Bot  `json:"bots"`
	Map      Map    `json:"map"`
	Goal     Goal   `json:"goal"`
	Queued   uint32 `json:"queued"`
	Started  uint32 `json:"started"`
	Finished uint32 `json:"finished"`
	Duration uint32 `json:"duration"`
}

type Goal struct {
	Age    uint32  `json:"age"`
	Ants   uint32  `json:"ants"`
	Hive   uint32  `json:"hive"`
	Alone  bool    `json:"alone"`
	Points []Point `json:"points"`
}
