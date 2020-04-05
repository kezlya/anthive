package anthive

type Game struct {
	ID       string `json:"id" firestore:"id"`
	Author   string `json:"author" firestore:"author"`
	Flavor   string `json:"flavor" firestore:"flavor"`
	Status   string `json:"status" firestore:"status"`
	Limit    uint16 `json:"limit" firestore:"limit"`
	Bots     []Bot  `json:"bots" firestore:"bots"`
	Map      `json:"map" firestore:"map"`
	Goal     `json:"goal" firestore:"goal"`
	Queued   int64 `json:"queued" firestore:"queued"`
	Started  int64 `json:"started" firestore:"started"`
	Finished int64 `json:"finished" firestore:"finished"`
	Duration int64 `json:"duration" firestore:"duration"`
}

type Goal struct {
	Age       int           `json:"age" firestore:"age"`
	Ants      int           `json:"ants" firestore:"ants"`
	Hive      int           `json:"hive" firestore:"hive"`
	Alone     bool          `json:"alon" firestore:"alon"`
	Positions []Coordinates `json:"position" firestore:"position"`
}
