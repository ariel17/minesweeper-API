package games

import "time"

const (
	// Game statuses
	startedStatus  = "started"
	youLooseStatus = "you_loose"
	youWonStatus   = "you_won"

	// Cell statuses
	coveredStatus = "covered"
	cleanedStatus = "cleaned"
	markedStatus  = "marked"
)

var (
	games []*Game
)

// Config holds the parameters on which the game will be created.
type Config struct {
	Rows    int64 `json:"rows"`
	Columns int64 `json:"columns"`
	Mines   int64 `json:"mines"`
}

// Cell is the representation of a single cell in the board.
// [0][1][2]
// [3][c][4]
// [5][6][7]
type Cell struct {
	IsMine    bool    `json:"is_mine"`
	Status    string  `json:"status"`
	Adjacents []*Cell `json:"adjacents"`
	Game      *Game   `json:"-"`
}

func newCell(game *Game, adjacents []*Cell) *Cell {
	return &Cell{
		Status:    coveredStatus,
		Adjacents: adjacents,
		Game:      game,
	}
}

// Clear opens the content of the cell.
func (c *Cell) Clear(caller *Cell) {
	if caller == nil { // from user
		c.Status = cleanedStatus
		if c.IsMine {
			c.Game.Blow()
			return
		}
		for _, a := range c.Adjacents {
			a.Clear(c)
		}
		// TODO make method
		// } else if c.IsAdjacent(caller) && !c.IsMine {
		// 	c.Status = cleanedStatus
	}
}

// Game is the representation of a playeable game.
type Game struct {
	ID         int64      `json:"id"`
	Config     Config     `json:"config"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	board      [][]*Cell  `json:"-"`
}

// New creates a new instance of the game.
func New() *Game {
	g := Game{
		ID: int64(len(games) + 1),
		Config: Config{
			Rows:    50,
			Columns: 50,
			Mines:   10,
		},
		Status:    startedStatus,
		CreatedAt: time.Now(),
	}
	g.build()
	games = append(games, &g)
	return &g
}

func (g *Game) build() {
	g.board = make([][]*Cell, g.Config.Rows)
	for i := range g.board {
		g.board[i] = make([]*Cell, g.Config.Columns)
		for j := range g.board[i] {
			adjacents := []*Cell{}
			g.board[i][j] = newCell(g, adjacents)
		}
	}
}

// Blow indicates that the user found a mine and losed the game.
func (g *Game) Blow() {
	g.Status = youLooseStatus
	g.finish()
}

// Clear indidicates that the user won the game.
func (g *Game) Clear() {
	g.Status = youWonStatus
	g.finish()
}

func (g *Game) finish() {
	finished := time.Now()
	g.FinishedAt = &finished
}
