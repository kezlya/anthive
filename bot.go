package anthive

type Bot struct {
	Lang       string `json:"lang"`
	Source     string `json:"source"`
	Version    string `json:"version"`
	Skin       uint   `json:"skin"`
	Conversion uint   `json:"conversion"`
	Span       Point  `json:"Point"`
	Username   string `json:"username"`
	Wealth     string `json:"wealth"`
	Stats      Stats  `json:"stats"`
	Avatar     string `json:"avatar"`
}

type Stats struct {
	Age    uint `json:"age"`
	Art    uint `json:"art"`
	Ants   uint `json:"ants"`
	Hive   uint `json:"hive"`
	Errors uint `json:"errors"`
}
