package anthive

type Request struct {
	Id     string `json:"id"`
	Tick   uint32 `json:"tick"`
	Ants   []Ant  `json:"ants"`
	Canvas Canvas `json:"canvas"`
}

type Response struct {
	Orders []Order `json:"orders"`
}

type Ant struct {
	Id      uint32 `json:"id"`
	Event   string `json:"event"`
	Errors  uint32 `json:"errors"`
	Age     uint32 `json:"age"`
	Health  uint32 `json:"health"`
	Payload uint32 `json:"payload"`
	Point   Point  `json:"point"`
}

type Order struct {
	AntId     uint32 `json:"antId"`
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
