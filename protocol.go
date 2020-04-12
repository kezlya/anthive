package anthive

type BotRequest struct {
	ID      string          `json:"id"`
	Tick    uint16          `json:"tick"`
	Ants    map[uint16]*Ant `json:"ants"`
	*Canvas `json:"canvas"`
}

type BotResponse struct {
	Orders       map[uint16]*Order `json:"orders"`
	Milliseconds int               `json:"milliseconds"`
}

type Ant struct {
	Event   string `json:"event"`
	Errors  int    `json:"errors"`
	Age     uint8  `json:"age"`
	Health  uint8  `json:"health"`
	Payload uint8  `json:"payload"`
	X       uint8  `json:"x"`
	Y       uint8  `json:"y"`
}

type Canvas struct {
	Width  uint8     `json:"width"`
	Height uint8     `json:"height"`
	Cells  [][]*Cell `json:"cells"`
}

type Order struct {
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
