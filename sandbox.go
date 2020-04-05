package anthive

type Sandbox struct {
	Code     string `json:"code" firestore:"code"`
	Game     `json:"game" firestore:"game"`
	Lang     string `json:"lang" firestore:"lang"`
	Results  `json:"results" firestore:"results"`
	Status   string `json:"status" firestore:"status"`
	Queued   int64  `json:"queued" firestore:"queued"`
	Duration int64  `json:"duration" firestore:"duration"`
}

type Results struct {
	Age      int    `json:"age" firestore:"age"`
	Logs     string `json:"request" firestore:"request"`
	Response string `json:"response" firestore:"response"`
	Data     string `json:"data" firestore:"data"`
}

type Data struct {
	Frames [][]CellDiff `json:"frames" firestore:"frames"`
}

type CellDiff struct {
	L uint16 `json:"l" firestore:"v"`
	V uint16 `json:"l" firestore:"v"`
}
