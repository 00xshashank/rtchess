package types

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Rating    uint   `json:"rating"`
}

type GameResult uint

const (
	ResultWhiteWon GameResult = iota
	ResultBlackWon
	ResultDraw
	ResultGameInProgress
)
