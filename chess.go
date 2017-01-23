package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	startGame()
	setup()
	fmt.Println("Starting game")
}

const (
	GRIDLENGTH    uint = 8
	playerFaction      = 0
)

type piece struct {
	faction uint
	class   uint
}

type position struct {
	x uint
	y uint
}

type move struct {
	from position
	to   position
}

type grid [GRIDLENGTH][GRIDLENGTH]*piece

func (move move) Value(g *grid) int {
	val := 0
	target := getGrid(g, move.to)
	if target != nil {
		switch target.class {
		case 0:
			val = 1
		case 1:
			val = 3
		case 2:
			val = 3
		case 3:
			val = 5
		case 4:
			val = 9
		case 5:
			val = 20
		}
	}
	return val
}

type moveSet struct {
	moves []move
	g     *grid
}

func (moveSet moveSet) Len() int {
	return len(moveSet.moves)
}
func (moveSet moveSet) Less(i, j int) bool {
	// i value
	ival := moveSet.moves[i].Value(moveSet.g)

	// j value
	jval := moveSet.moves[j].Value(moveSet.g)

	return ival > jval
}
func (moveSet moveSet) Swap(i, j int) {
	temp := moveSet.moves[i]
	moveSet.moves[i] = moveSet.moves[j]
	moveSet.moves[j] = temp
}

var (
	board    *grid
	selected *position
)

// User function
func grabPiece(x uint, y uint) {
	from := position{x: x, y: y}
	piece := getGrid(board, from)
	if piece != nil && piece.faction == playerFaction {
		selected = &from
	}
}

// User function
func movePiece(x uint, y uint) {
	to := position{x: x, y: y}
	if selected != nil {
		currentMove := move{from: *selected, to: to}
		if validateMove(board, currentMove, 0) {
			executeMove(board, currentMove)
			selected = nil
			enemyMove()
		}
	}
}

func executeMove(g *grid, m move) {
	setGrid(g, m.to, getGrid(g, m.from))
	setGrid(g, m.from, nil)
}

func setGrid(g *grid, pos position, piece *piece) {
	g[pos.x][pos.y] = piece
}

func getGrid(g *grid, pos position) *piece {
	return g[pos.x][pos.y]
}

func enemyMove() {
	moveSet := findAllValidMoves(board, 1, 0)
	if len(moveSet.moves) == 0 {
		os.Exit(0)
	}
	sort.Sort(moveSet)

	choices := len(moveSet.moves)
	if choices > 5 {
		choices = 5
	}

	chosen := moveSet.moves[rand.Intn(choices)]
	executeMove(board, chosen)
}

func findAllValidMoves(g *grid, faction uint, level uint) moveSet {

	validMoves := [1024]move{}
	moveCount := uint(0)

	for fx := uint(0); fx < GRIDLENGTH; fx++ {
		for fy := uint(0); fy < GRIDLENGTH; fy++ {
			from := position{x: fx, y: fy}
			fromPiece := getGrid(g, from)
			if fromPiece != nil && fromPiece.faction == faction {
				for x := uint(0); x < GRIDLENGTH; x++ {
					for y := uint(0); y < GRIDLENGTH; y++ {
						to := position{x: x, y: y}
						move := move{from: from, to: to}
						if validateMove(g, move, level) {
							validMoves[moveCount] = move
							moveCount++
						}
					}
				}
			}
		}
	}

	return moveSet{g: g, moves: validMoves[:moveCount]}
}

func validateMove(g *grid, move move, level uint) bool {
	movingPiece := getGrid(g, move.from)
	if movingPiece == nil {
		return false
	}
	movingPieceFaction := movingPiece.faction
	oppositeFaction := (movingPieceFaction + 1) % 2

	// validate the movement
	movement := validateMovement(g, move)

	if !movement {
		return false
	}

	if level == 0 {
		isNotIntoCheck := isNotIntoCheck(g, move, oppositeFaction)

		if !isNotIntoCheck {
			return false
		}
	}

	return true
}

func isNotIntoCheck(g *grid, move move, oppositeFaction uint) bool {
	// validate whether or not the movement puts the mover in check
	// clone the grid and do the move on the hypergrid
	// then check whether or not the current player will be checkmate next turn
	var hypergrid grid
	hypergrid = *g
	hypergridpointer := &hypergrid

	executeMove(hypergridpointer, move)

	validOpponentMoves := findAllValidMoves(hypergridpointer, oppositeFaction, 1)
	// ensure that a valid move for the oponent isn't checkmating the current player
	for _, move := range validOpponentMoves.moves {
		if move.Value(hypergridpointer) == 20 {
			return false
		}
	}

	return true
}

func isCheckMate(faction uint) bool {
	return false
}

// from and to
func validateMovement(g *grid, move move) bool {
	fx := move.from.x
	fy := move.from.y
	x := move.to.x
	y := move.to.y

	// selected piece
	selectedPiece := g[fx][fy]
	// target piece
	target := g[x][y]

	switch selectedPiece.class {
	// pawn
	case 0:
		var canDoubleStep bool
		dy := int(y) - int(fy)
		dx := int(fx) - int(x)
		var advance int
		if selectedPiece.faction == 0 {
			// white
			advance = -dy
			canDoubleStep = fy == 6 && g[x][5] == nil
		} else {
			// black
			advance = dy
			canDoubleStep = fy == 1 && g[x][2] == nil
		}

		if fx == x {
			return target == nil && (advance == 1 || (canDoubleStep && advance == 2))
		} else {
			sidemovement := dx
			if sidemovement < 0 {
				sidemovement = -sidemovement
			}

			return advance == 1 && sidemovement == 1 && target != nil && target.faction != selectedPiece.faction
		}
	// knight
	case 1:
		dx := int(x) - int(fx)
		if dx < 0 {
			dx = -dx
		}
		dy := int(y) - int(fy)
		if dy < 0 {
			dy = -dy
		}

		if (dx == 2 && dy == 1) || (dy == 2 && dx == 1) {
			if target == nil || (target.faction != selectedPiece.faction) {
				return true
			}
		}
		return false
	// bishop
	case 2:
		dx := int(x) - int(fx)
		vecx := dx
		if dx < 0 {
			dx = -dx
		}

		dy := int(y) - int(fy)
		vecy := dy
		if dy < 0 {
			dy = -dy
		}

		if dx != dy || dx == 0 {
			return false
		}

		incx := vecx / dx
		incy := vecy / dy

		for tx, ty := int(fx)+incx, int(fy)+incy; tx != int(x); tx, ty = tx+incx, ty+incy {
			if g[tx][ty] != nil {
				return false
			}
		}
		return target == nil || target.faction != selectedPiece.faction
	// rook
	case 3:
		dx := int(x) - int(fx)
		vecx := dx
		if dx < 0 {
			dx = -dx
		}

		dy := int(y) - int(fy)
		vecy := dy
		if dy < 0 {
			dy = -dy
		}

		// check movement is only in one dimension
		if (dx != 0) == (dy != 0) {
			return false
		}

		incx := 0
		if dx != 0 {
			incx = vecx / dx
		}
		incy := 0
		if dy != 0 {
			incy = vecy / dy
		}

		for tx, ty := int(fx)+incx, int(fy)+incy; tx != int(x) || ty != int(y); tx, ty = tx+incx, ty+incy {
			if g[tx][ty] != nil {
				return false
			}
		}
		return target == nil || target.faction != selectedPiece.faction
	// queen
	case 4:
		dx := int(x) - int(fx)
		vecx := dx
		if dx < 0 {
			dx = -dx
		}

		dy := int(y) - int(fy)
		vecy := dy
		if dy < 0 {
			dy = -dy
		}

		// check movement is valid
		if dx != dy && ((dx != 0) == (dy != 0)) {
			return false
		}

		incx := 0
		if dx != 0 {
			incx = vecx / dx
		}
		incy := 0
		if dy != 0 {
			incy = vecy / dy
		}

		for tx, ty := int(fx)+incx, int(fy)+incy; tx != int(x) || ty != int(y); tx, ty = tx+incx, ty+incy {
			if g[tx][ty] != nil {
				return false
			}
		}
		return target == nil || target.faction != selectedPiece.faction
	// king
	case 5:
		dx := int(x) - int(fx)
		vecx := dx
		if dx < 0 {
			dx = -dx
		}

		dy := int(y) - int(fy)
		vecy := dy
		if dy < 0 {
			dy = -dy
		}

		// check movement is valid
		if dx > 1 || dy > 1 {
			return false
		}

		incx := 0
		if dx != 0 {
			incx = vecx / dx
		}
		incy := 0
		if dy != 0 {
			incy = vecy / dy
		}

		for tx, ty := int(fx)+incx, int(fy)+incy; tx != int(x) || ty != int(y); tx, ty = tx+incx, ty+incy {
			if g[tx][ty] != nil {
				return false
			}
		}
		return target == nil || target.faction != selectedPiece.faction
	}

	return false
}

func startGame() {

	// 0 = white
	// 1 = black

	// 0 = pawn
	// 1 = knight
	// 2 = bishop
	// 3 = rook
	// 4 = queen
	// 5 = king

	// set up board
	var cb grid
	board = &cb

	// white pawns
	for x := uint(0); x < GRIDLENGTH; x++ {
		board[x][6] = &(piece{
			class: 0, faction: 0,
		})
	}

	board[0][7] = &(piece{
		class: 3, faction: 0,
	})

	board[1][7] = &(piece{
		class: 1, faction: 0,
	})

	board[2][7] = &(piece{
		class: 2, faction: 0,
	})

	board[3][7] = &(piece{
		class: 4, faction: 0,
	})

	board[4][7] = &(piece{
		class: 5, faction: 0,
	})

	board[5][7] = &(piece{
		class: 2, faction: 0,
	})

	board[6][7] = &(piece{
		class: 1, faction: 0,
	})

	board[7][7] = &(piece{
		class: 3, faction: 0,
	})

	// black pawns
	for x := uint(0); x < GRIDLENGTH; x++ {
		board[x][1] = &(piece{
			class: 0, faction: 1,
		})
	}

	board[0][0] = &(piece{
		class: 3, faction: 1,
	})

	board[1][0] = &(piece{
		class: 1, faction: 1,
	})

	board[2][0] = &(piece{
		class: 2, faction: 1,
	})

	board[3][0] = &(piece{
		class: 4, faction: 1,
	})

	board[4][0] = &(piece{
		class: 5, faction: 1,
	})

	board[5][0] = &(piece{
		class: 2, faction: 1,
	})

	board[6][0] = &(piece{
		class: 1, faction: 1,
	})

	board[7][0] = &(piece{
		class: 3, faction: 1,
	})

}
