package anthive

type Bot struct {
	Lang string `json:"lang"`
	//Source     string `json:"source"`
	Version    string `json:"version"`
	Skin       int    `json:"skin"`
	Conversion int    `json:"conversion"`
	Spawn      Point  `json:"spawn"`
	Username   string `json:"username"`
	Stats      Stats  `json:"stats"`
	Avatar     string `json:"avatar"`
}

type Stats struct {
	Score  string `json:"score"`
	Age    int    `json:"age"`
	Art    int    `json:"art"`
	Ants   int    `json:"ants"`
	Hive   int    `json:"hive"`
	Errors int    `json:"errors"`
}
