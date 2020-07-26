package anthive

type Bot struct {
	Lang       string `json:"lang"`
	Source     string `json:"source"`
	Version    string `json:"version"`
	Skin       uint32 `json:"skin"`
	Conversion uint32 `json:"conversion"`
	Span       Point  `json:"Point"`
	Username   string `json:"username"`
	Wealth     string `json:"wealth"`
	Stats      Stats  `json:"stats"`
	Avatar     string `json:"avatar"`
}

type Stats struct {
	Age    uint32 `json:"age"`
	Art    uint32 `json:"art"`
	Ants   uint32 `json:"ants"`
	Hive   uint32 `json:"hive"`
	Errors uint32 `json:"errors"`
}
