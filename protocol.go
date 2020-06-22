package anthive

type Request struct {
	Id     string `json:"id"`
	Tick   uint   `json:"tick"`
	Ants   []Ant  `json:"ants"`
	Canvas Canvas `json:"canvas"`
}

type Response struct {
	Orders []Order `json:"orders"`
}

type Ant struct {
	Id      uint   `json:"id"`
	Event   string `json:"event"`
	Errors  int    `json:"errors"`
	Age     uint   `json:"age"`
	Health  uint   `json:"health"`
	Payload uint   `json:"payload"`
	Point   Point  `json:"point"`
}

type Order struct {
	AntId     string `json:"antId"`
	Action    string `json:"act"`
	Direction string `json:"dir"`
}

const ActionStay = "stay"
const ActionMove = "move"
const ActionLoad = "load"
const ActionUnload = "unload"
const ActionEat = "eat"

const DirectionUp = "up"
const DirectionRight = "right"
const DirectionDown = "down"
const DirectionLeft = "left"

const EventGood = "good"
const EventBirth = "birth"
const EventNoAction = "noaction"
const EventSlow = "slow"
const EventBadMove = "badmove"
const EventBadLoad = "badload"
const EventBadUnload = "badunload"
const EventBadEat = "badeat"
const EventCollision = "collision"
const EventError = "error"
const EventDeath = "death"
